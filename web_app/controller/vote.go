package controller

import (
	"golearn/web_app/logic"
	"golearn/web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//投票

func PostVoteHandler(c *gin.Context) {

	//参数校验
	p := new(models.VoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		//
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvaildParam, err.Error())
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvaildParam, errData)
		return
	}

	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeInvaildLogin)
		return
	}
	//具体投票的业务逻辑
	if err := logic.PostVote(userID, p); err != nil {
		zap.L().Error("logic.PostVote failed", zap.Error(err))
		ResponseError(c, CodeInvaildLogin)
		return
	}
	ResponseSuccess(c, nil)
}
