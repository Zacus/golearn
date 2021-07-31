package controller

import (
	"golearn/web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//community处理社区请求的函数
func CommunityHandler(c *gin.Context) {

	//查询到所有的社区(community_id,community_name)以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)

}

//社区分区详情
func CommunityDetailHandler(c *gin.Context) {

	//获取社区id
	strCommunityID := c.Param("id")
	// id, err := strconv.ParseInt(strCommunityID, 10, 64)
	// if err != nil {
	// 	ResposeError(c, CodeInvaildParam)
	// 	return
	// }
	data, err := logic.GetCommunityDetailById(strCommunityID)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}
