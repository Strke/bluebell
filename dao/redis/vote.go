package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math"
	"time"
)

/*
投一票加432分
投票的集中清苦那个：
direction=1时，有两种情况：
	1、之前没投过票，现在投赞成票   +432
	2、之前投反对票，现在改投赞成票 +432*2
direction=0时，有两种情况：
	1、之前头赞成票，现在要取消投票 -432
	2、之前投反对票，现在要取消投票 +432
direction=-1时，有两种情况：
	1、之前没有投过票，现在投反对票 -432
	2、之前投赞成票，现在改投反对票 -432*2
投票限制：
每个帖子超过一星期就不让投票
*/

const (
	oneWeekInSecond = 7 * 24 * 3600
	scorePerVote    = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepested   = errors.New("不允许重复投票")
)

func CreatePost(postID int64) error {
	pipeline := client.TxPipeline()
	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	//1、判断投票的限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime < oneWeekInSecond {
		zap.L().Error("投票时间到达",
			zap.Float64("time-diff", float64(time.Now().Unix())-postTime-oneWeekInSecond))
		return ErrVoteTimeExpire
	}
	//2、更新帖子的分数
	//先查之前的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedPrefix+postID), userID).Val()
	if value == ov {
		return ErrVoteRepested
	}
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value)
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID)

	//3、记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
