package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"golearn/web_app/models"
)

const secret = "zs"

//把每一步数据库操作封装成数据库
//待logic层根据业务需求调用
func IsExist(name string) (err error) {

	sqlStr := "SELECT COUNT(user_id) FROM user WHERE username=?"
	var count int
	if err = db.Get(&count, sqlStr, name); err != nil {
		return
	}
	if count > 0 {
		err = nil
	}
	return

}

func InsertUser(user *models.User) (err error) {

	//对密码进行加密
	password := encryptPassword(user.Password)
	//执行SQL语句
	sqlStr := "insert into user(user_id,username,password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)
	return
}

func Login(user *models.User) (err error) {

	oPassword := user.Password
	sqlStr := "select user_id,username,password from user where username=?"
	if err = db.Get(user, sqlStr, user.UserName); err != nil {
		return
	}
	// if err == sql.ErrNoRows {
	// 	return errors.New("用户不存在")
	// }

	//对密码进行加密
	oPassword = encryptPassword(oPassword)

	if oPassword != user.Password {
		return errors.New("密码错误")
	}
	return

}

func encryptPassword(oPassword string) string {

	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
