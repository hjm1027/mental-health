package model

import "time"

type CourseModel struct {
	Id          uint32    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Url         string    `gorm:"column:url"`
	Name        string    `gorm:"column:name"`
	Source      string    `gorm:"column:source"`
	Summary     string    `gorm:"column:summary"`
	LikeNum     uint32    `gorm:"column:like_num"`
	FavoriteNum uint32    `gorm:"column:favorite_num"`
	WatchNum    uint32    `gorm:"column:watch_num"`
	Time        time.Time `gorm:"column:time"`
}

type CourseInfoResponse struct {
	Id          uint32    `json:"id"`
	Url         string    `json:"url"`
	Name        string    `json:"name"`
	Source      string    `json:"source"`
	Summary     string    `json:"summary"`
	LikeNum     uint32    `json:"like_num"`
	FavoriteNum uint32    `json:"favorite_num"`
	WatchNum    uint32    `json:"watch_num"`
	Time        time.Time `json:"time"`
	IsLike      bool      `json:"is_like"`
	IsFavorite  bool      `json:"is_favorite"`
}

// 课程点赞中间表
type CourseLikeModel struct {
	Id       uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserId   uint32 `gorm:"column:user_id"`
	CourseId uint32 `gorm:"column:course_id"`
}

// 课程收藏中间表
type CourseFavoriteModel struct {
	Id       uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserId   uint32 `gorm:"column:user_id"`
	CourseId uint32 `gorm:"column:course_id"`
}
