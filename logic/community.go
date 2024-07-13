package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// 返回社区信息（community_id, community_name）

func GetCommunityList() ([]*models.Community, error) {
	//查数据库
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
