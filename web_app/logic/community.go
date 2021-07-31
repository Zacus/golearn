package logic

import (
	"golearn/web_app/dao/mysql"
)

func GetCommunityList() (data interface{}, err error) {
	//查数据库，查找到所有的community, 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetailById(idStr string) (data interface{}, err error) {

	return mysql.GetCommunityByID(idStr)
}
