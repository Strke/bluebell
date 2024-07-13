package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	//1、查询到所有的社区（community_id, community_name）
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("Logic.GetCommunityList() Failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	//1、获取社区id
	idStr := c.Param("id")
	//2、参数校验
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//1、查询到所有的社区（community_id, community_name）
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("Logic.GetCommunityDetail() Failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
