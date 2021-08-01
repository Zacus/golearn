package redis

/*
	Redis Key
*/

const (
	KeyPostInfoHashPrefix = "bluebell:post:"
	KeyPostTimeZSet       = "bluebell:post:time"
	KeyPostScoreZSet      = "bluebell:post:score"
	//KeyPostVotedUpSetPrefix   = "bluebell:post:voted:down:"
	//KeyPostVotedDownSetPrefix = "bluebell:post:voted:up:"
	KeyPostVotedZSetPrefix = "bluebell:post:voted:" //zset;记录用户及投票类型；参数是post_id

	KeyCommunityPostSetPrefix = "bluebell:community:"
)
