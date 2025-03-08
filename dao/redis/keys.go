package redis

//redis key尽量使用命名空间方式，区分不同的key，方便业务拆分和查询

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "bluebell:post:time"  //zset;<发帖时间：帖子ID>
	KeyPostScoreZSet   = "bluebell:post:score" //zset;<帖子得分：帖子ID>
	KeyPostVotedPrefix = "bluebell:post:voted:" //zset;帖子ID：<用户ID：投票类型>
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
