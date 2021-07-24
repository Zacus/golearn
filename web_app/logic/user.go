package logic

import (
	"golearn/web_app/dao/mysql"
	"golearn/web_app/models"
	"golearn/web_app/pkg/jwt"
	"golearn/web_app/pkg/snowflake"

	"go.uber.org/zap"
)

//存放业务逻辑的代码
//注册逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	err = mysql.IsExist(p.UserName)
	if err != nil {
		zap.L().Error("query failed", zap.Error(err))
		return
	}

	//生成user_id
	userid := snowflake.GenID()

	//创建User实例
	u := &models.User{
		UserID:   userid,
		UserName: p.UserName,
		Password: p.Password,
	}
	//password加密
	//存入数据库
	if err = mysql.InsertUser(u); err != nil {
		zap.L().Error("database insert  failed", zap.Error(err))
		return
	}
	return

}

//登录逻辑
func Login(p *models.ParamLogin) (token string, err error) {
	//判断用户是否存在，判断密码是否正确
	token = ""
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		zap.L().Error("Login  failed", zap.Error(err))
		return
	}

	//生成token
	return jwt.GenToken(user.UserID, user.UserName)

}
