package redis

//redis key尽量使用命名空间方式，区分不同的key，方便业务拆分和查询

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "bluebell:post:time"   //zset;帖子及发帖时间为分数
	KeyPostScoreZSet   = "bluebell:post:score"  //zset;帖子及投票的分数
	KeyPostVotedPrefix = "bluebell:post:voted:" //zset;记录用户及投票类型，参数是post id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
