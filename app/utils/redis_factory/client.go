package redis_factory

import (
	"douyin-backend/app/core/event_manage"
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/utils/yml_config"
	"douyin-backend/app/utils/yml_config/interf"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

var redisPool *redis.Pool
var configYml interf.YmlConfigInterface

// 处于程序底层的包，init 初始化的代码段的执行会优先于上层代码，因此这里读取配置项不能使用全局配置项变量
func init() {
	configYml = yml_config.CreateYamlFactory()
	redisPool = initRedisClientPool()
}
func initRedisClientPool() *redis.Pool {
	redisPool = &redis.Pool{
		MaxIdle:     configYml.GetInt("Redis.MaxIdle"),                        //最大空闲数
		MaxActive:   configYml.GetInt("Redis.MaxActive"),                      //最大活跃数
		IdleTimeout: configYml.GetDuration("Redis.IdleTimeout") * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			//此处对应redis ip及端口号
			conn, err := redis.Dial("tcp", configYml.GetString("Redis.Host")+":"+configYml.GetString("Redis.Port"))
			if err != nil {
				variable.ZapLog.Error(my_errors.ErrorsRedisInitConnFail + err.Error())
				return nil, err
			}
			auth := configYml.GetString("Redis.Auth") //通过配置项设置redis密码
			if len(auth) >= 1 {
				if _, err := conn.Do("AUTH", auth); err != nil {
					_ = conn.Close()
					variable.ZapLog.Error(my_errors.ErrorsRedisAuthFail + err.Error())
				}
			}
			_, _ = conn.Do("select", configYml.GetInt("Redis.IndexDb"))
			return conn, err
		},
	}
	// 将redis的关闭事件，注册在全局事件统一管理器，由程序退出时统一销毁
	eventManageFactory := event_manage.CreateEventManageFactory()
	if _, exists := eventManageFactory.Get(variable.EventDestroyPrefix + "Redis"); exists == false {
		eventManageFactory.Set(variable.EventDestroyPrefix+"Redis", func(args ...interface{}) {
			_ = redisPool.Close()
		})
	}
	return redisPool
}

// 从连接池获取一个redis连接
func GetOneRedisClient() *RedisClient {
	maxRetryTimes := configYml.GetInt("Redis.ConnFailRetryTimes")
	var oneConn redis.Conn
	for i := 1; i <= maxRetryTimes; i++ {
		oneConn = redisPool.Get()
		// 首先通过执行一个获取时间的命令检测连接是否有效，如果已有的连接无法执行命令，则重新尝试连接到redis服务器获取新的连接池地址
		// 连接不可用可能会发生的场景主要有：服务端redis重启、客户端网络在有线和无线之间切换等
		if _, replyErr := oneConn.Do("time"); replyErr != nil {
			//fmt.Printf("连接已经失效(出错)：%+v\n", replyErr.Error())
			// 如果已有的redis连接池获取连接出错(官方库的说法是连接不可用)，那么继续使用从新初始化连接池
			initRedisClientPool()
			oneConn = redisPool.Get()
		}
		if oneConn.Err() != nil {
			//variable.ZapLog.Error("Redis：网络中断,开始重连进行中..." , zap.Error(oneConn.Err()))
			if i == maxRetryTimes {
				variable.ZapLog.Error(my_errors.ErrorsRedisGetConnFail, zap.Error(oneConn.Err()))
				return nil
			}
			//如果出现网络短暂的抖动，短暂休眠后，支持自动重连
			time.Sleep(time.Second * configYml.GetDuration("Redis.ReConnectInterval"))
		} else {
			break
		}
	}
	return &RedisClient{oneConn}
}

// 定义一个redis客户端结构体
type RedisClient struct {
	client redis.Conn
}

// 为redis-go 客户端封装统一操作函数入口
func (r *RedisClient) Execute(cmd string, args ...interface{}) (interface{}, error) {
	return r.client.Do(cmd, args...)
}

// 释放连接到连接池
func (r *RedisClient) ReleaseOneRedisClient() {
	_ = r.client.Close()
}

//  封装几个数据类型转换的函数

// bool 类型转换
func (r *RedisClient) Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}

// string 类型转换
func (r *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

// strings 类型转换
func (r *RedisClient) Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

// Float64 类型转换
func (r *RedisClient) Float64(reply interface{}, err error) (float64, error) {
	return redis.Float64(reply, err)
}

// int 类型转换
func (r *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

// int64 类型转换
func (r *RedisClient) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

// uint64 类型转换
func (r *RedisClient) Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}

// Bytes 类型转换
func (r *RedisClient) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}
