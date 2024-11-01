package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/video"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoController struct {
}

func (v *VideoController) VideoDigg(ctx *gin.Context) {
	var uid = ctx.GetString(consts.ValidatorPrefix + "uid")
	var aweme_id = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var uidInt64, _ = strconv.ParseInt(uid, 10, 64)
	var awemeIDInt64, _ = strconv.ParseInt(aweme_id, 10, 64)
	diggDone := video.CreateDiggFactory("").VideoDigg(uidInt64, awemeIDInt64)
	if diggDone {
		ctx.JSON(http.StatusOK, gin.H{
			"data": diggDone,
			"code": consts.CurdStatusOkCode,
			"msg":  "点赞成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"data": diggDone,
			"code": consts.CurdCreatFailCode,
			"msg":  "点赞失败",
		})
	}
}

func (v *VideoController) VideoComment(ctx *gin.Context) {
	var ip_location = ctx.GetString(consts.ValidatorPrefix + "ip_location")
	var aweme_id = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var content = ctx.GetString(consts.ValidatorPrefix + "content")
	var uid = ctx.GetString(consts.ValidatorPrefix + "uid")
	var short_id = ctx.GetString(consts.ValidatorPrefix + "short_id")
	var unique_id = ctx.GetString(consts.ValidatorPrefix + "unique_id")
	var signature = ctx.GetString(consts.ValidatorPrefix + "signature")
	var nickname = ctx.GetString(consts.ValidatorPrefix + "nickname")
	var avatar = ctx.GetString(consts.ValidatorPrefix + "avatar")
	var uidInt64, _ = strconv.ParseInt(uid, 10, 64)
	var awemeIDInt64, _ = strconv.ParseInt(aweme_id, 10, 64)
	commentDone := video.CreateCommentFactory("").VideoComment(uidInt64, awemeIDInt64, ip_location, content, short_id, unique_id, signature, nickname, avatar)
	if commentDone {
		ctx.JSON(http.StatusOK, gin.H{
			"data": commentDone,
			"code": consts.CurdStatusOkCode,
			"msg":  "评论成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"data": commentDone,
			"code": consts.CurdCreatFailCode,
			"msg":  "评论失败",
		})
	}

}

func (v *VideoController) VideoCollect(ctx *gin.Context) {
	var uid = ctx.GetString(consts.ValidatorPrefix + "uid")
	var aweme_id = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var uidInt64, _ = strconv.ParseInt(uid, 10, 64)
	var awemeIDInt64, _ = strconv.ParseInt(aweme_id, 10, 64)
	diggDone := video.CreateCollectFactory("").VideoCollect(uidInt64, awemeIDInt64)
	if diggDone {
		ctx.JSON(http.StatusOK, gin.H{
			"data": diggDone,
			"code": consts.CurdStatusOkCode,
			"msg":  "收藏成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"data": diggDone,
			"code": consts.CurdCreatFailCode,
			"msg":  "收藏失败",
		})
	}
}

func (v *VideoController) VideoShare(ctx *gin.Context) {
	var uid = ctx.GetString(consts.ValidatorPrefix + "uid")
	var aweme_id = ctx.GetString(consts.ValidatorPrefix + "aweme_id")
	var message = ctx.GetString(consts.ValidatorPrefix + "message")
	var share_uid_list = ctx.GetString(consts.ValidatorPrefix + "share_uid_list")
	var uidInt64, _ = strconv.ParseInt(uid, 10, 64)
	var awemeIDInt64, _ = strconv.ParseInt(aweme_id, 10, 64)
	fmt.Println(share_uid_list)
	shareDone := video.CreateShareFactory("").VideoShare(uidInt64, awemeIDInt64, message, share_uid_list)
	if shareDone {
		ctx.JSON(http.StatusOK, gin.H{
			"data": shareDone,
			"code": consts.CurdStatusOkCode,
			"msg":  "分享成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"data": shareDone,
			"code": consts.CurdCreatFailCode,
			"msg":  "分享失败",
		})
	}
}

func (v *VideoController) GetComments(ctx *gin.Context) {
	aweme_id, _ := strconv.Atoi(ctx.Query("aweme_id"))
	comments := video.CreateCommentFactory("").GetComments(int64(aweme_id))
	if len(comments) > 0 {
		ctx.JSON(http.StatusOK, comments)
	} else {
		ctx.JSON(http.StatusNoContent, comments)
	}
}

func (v *VideoController) GetStar(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (v *VideoController) GetShare(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (v *VideoController) GetHistoryOther(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}

func (v *VideoController) GetLongVideoRecommended(ctx *gin.Context) {
	// TODO 具体业务逻辑实现
	var Uid = ctx.GetFloat64(consts.ValidatorPrefix + "uid")
	var PageNo = ctx.GetFloat64(consts.ValidatorPrefix + "pageNo")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := video.CreateVideoFactory("").GetLongVideoRecommended(int64(Uid), int64(PageNo), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": total,
			"list":  list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
			"total": total,
			"list":  []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (v *VideoController) GetVideoRecommended(ctx *gin.Context) {
	Uid, _ := strconv.Atoi(ctx.Query("uid"))
	var Start = ctx.GetFloat64(consts.ValidatorPrefix + "start")
	var PageSize = ctx.GetFloat64(consts.ValidatorPrefix + "pageSize")
	list, total := video.CreateVideoFactory("").GetVideoRecommended(int64(Uid), int64(Start), int64(PageSize))
	if len(list) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": total,
			"list":  list,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{
			"total": total,
			"list":  []interface{}{}, // 返回一个空数组以确保响应一致性
		})
	}
}

func (v *VideoController) GetHistory(context *gin.Context) {
	// TODO 具体业务逻辑实现

	//var id = context.GetFloat64(consts.ValidatorPrefix + "id")
	//video := sv_home.CreateShortVideoFactory("").GetVideoById(int(id))
	//if video.Id != 0 {
	//	response.Success(context, consts.CurdStatusOkMsg, video)
	//} else {
	//	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	//}
}
