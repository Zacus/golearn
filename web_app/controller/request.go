package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID"

var ErrorUsernotLogin = errors.New("用户未登录")

/*
	func:getCurrentUser
	param:
	Description：获取当前登录的用户id

*/
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUsernotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUsernotLogin
		return
	}
	return
}
