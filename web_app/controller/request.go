package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID"

var ErrorUsernotLogin = errors.New("用户未登录")

/*
	func:getCurrentUserID
	param:
	Description：获取当前登录的用户id

*/
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	fmt.Println(ok)
	fmt.Println(uid)
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
