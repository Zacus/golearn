package logic

import (
	"golearn/web_app/dao/redis"
	"golearn/web_app/models"
	"strconv"

	"go.uber.org/zap"
)

//投票功能
/*
	1.用户投票的数据
	2.


*/

func PostVote(userID int64, p *models.VoteData) error {

	zap.L().Debug("PostVote",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.PostVote(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
	//1.判断投票限制

	//2.更新帖子的分数
	//3.记录用户为该帖子投票的数据

}
