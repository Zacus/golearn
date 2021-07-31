package controller

import (
	"fmt"
	"golearn/web_app/logic"
	"golearn/web_app/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// PostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {

	//获取参数
	post := new(models.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		ResponseErrorWithMsg(c, CodeInvaildParam, err.Error())
		return
	}
	// 参数校验

	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Debug("GetCurrentUserID() failed", zap.Any("err", err))
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeInvaildLogin)
		return
	}
	post.AuthorId = userID

	//2.创造帖子
	err = logic.CreatePost(post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// PostListHandler 帖子列表
// func PostListHandler(c *gin.Context) {
// 	//获取数据
// 	order, _ := c.GetQuery("order")
// 	pageStr, ok := c.GetQuery("page")
// 	if !ok {
// 		pageStr = "1"
// 	}
// 	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
// 	if err != nil {
// 		pageNum = 1
// 	}
// 	posts := redis.GetPost(order, pageNum)
// 	fmt.Println(len(posts))
// 	ResponseSuccess(c, posts)
// }

func PostListHandler(c *gin.Context) {
	//获取数据
	strPageStart, ok := c.GetQuery("page_start")
	if !ok {
		strPageStart = "1"
	}
	strPageLimit, _ := c.GetQuery("page_limit")

	page_start, err := strconv.ParseInt(strPageStart, 10, 64)
	if err != nil {
		page_start = 1
	}

	page_limit, err := strconv.ParseInt(strPageLimit, 10, 64)
	if err != nil {
		page_limit = 10
	}

	posts, err := logic.GetPostList(page_start, page_limit)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Println(len(posts))
	ResponseSuccess(c, posts)

}

func PostList2Handler(c *gin.Context) {
	data, err := logic.GetPostList2()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

// PostDetailHandler 帖子详情
func PostDetailHandler(c *gin.Context) {
	//1.获取参数(从url中获取帖子的id)
	postId := c.Param("id")

	//根据id取出帖子数据
	post, err := logic.GetPostByID(postId)
	if err != nil {
		zap.L().Error("logic.GetPost(postID) failed", zap.String("postId", postId), zap.Error(err))
		return
	}

	ResponseSuccess(c, post)
}
