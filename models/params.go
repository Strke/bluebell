package models

// 定义请求的参数、结构体

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 投票

type ParamVoteData struct {
	//UserID 从请求中获取当前用户
	PostID    string `json:"post_id" binding:"required"`              //帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票(1)还是反对票(-1),取消投票（0）
}

type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamCommunityPostList struct {
	ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
