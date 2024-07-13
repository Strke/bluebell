package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	//1、获取参数以及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldbindJson(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c中获取当前发请求的用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2、创建帖子查询数据库
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed!", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3、返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的函数
func GetPostDetailHandler(c *gin.Context) {
	//1、获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2、根据id取出帖子的数据
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("loggic.GetPostByID(pid) failed! get data from databases error!", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3、返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler获取帖子列表
func GetPostListHandler(c *gin.Context) {
	//获取分页参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	//获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// 根据前端传过来的排序属性（分数、创建时间）进行排序获取帖子列表
/*
1、获取参数
2、去redis查询id
3、根据id去数据库查询帖子详细信息
*/

func GetPostListHandler2(c *gin.Context) {
	//获取分页参数
	// api请求参数 /api/v1/posts2?page=1&size=10&order=time

	//pageStr := c.Query("page")
	//sizeStr := c.Query("size")
	//
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		return
	}

	//var (
	//	page int64
	//	size int64
	//	err  error
	//)
	//page, err = strconv.ParseInt(pageStr, 10, 64)
	//if err != nil {
	//	page = 1
	//}
	//size, err = strconv.ParseInt(sizeStr, 10, 64)
	//if err != nil {
	//	size = 10
	//}

	//获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}
