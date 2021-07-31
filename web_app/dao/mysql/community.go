package mysql

import (
	"database/sql"
	"golearn/web_app/models"

	"go.uber.org/zap"
)

//获取社区所有数据
func GetCommunityList() (communityList []*models.Community, err error) {

	sqlStr := "SELECT community_id,community_name FROM community"
	//查询数据库
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("There is no community data in the database.")
			err = nil
		}
	}
	return

}

// func GetCommunityDetailById(id int64) (community *models.Community, err error) {

// }

func GetCommunityNameByID(idStr string) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := `select community_id, community_name
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

func GetCommunityByID(idStr string) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}
