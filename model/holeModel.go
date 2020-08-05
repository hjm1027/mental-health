package model

import "time"

type HoleModel struct {
	Id          uint32    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserId      uint32    `gorm:"column:user_id"`
	Name        string    `gorm:"column:name"`
	Content     string    `gorm:"column:content"`
	LikeNum     uint32    `gorm:"column:like_num"`
	FavoriteNum uint32    `gorm:"column:favorite_num"`
	CommentNum  uint32    `gorm:"column:comment_num"`
	ReadNum     uint32    `gorm:"column:read_num"`
	Type        uint8     `gorm:"column:type"`
	Time        time.Time `gorm:"column:time"`
}

type HoleLikeModel struct {
	Id     uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	HoleId uint32 `gorm:"column:hole_id"`
	UserId uint32 `gorm:"column:user_id"`
}

type HoleFavoriteModel struct {
	Id     uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	HoleId uint32 `gorm:"column:hole_id"`
	UserId uint32 `gorm:"column:user_id"`
}

type HoleReadModel struct {
	Id     uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	HoleId uint32 `gorm:"column:hole_id"`
	UserId uint32 `gorm:"column:user_id"`
}

type HoleInfoResponse struct {
	HoleId      uint32           `json:"hole_id"`
	Type        uint8            `json:"type"`
	Content     string           `json:"content"`
	LikeNum     uint32           `json:"like_num"`
	ReadNum     uint32           `json:"read_num"`
	FavoriteNum uint32           `json:"favorite_num"`
	IsLike      bool             `json:"is_like"`
	IsFavorite  bool             `json:"is_favorite"`
	Time        time.Time        `json:"time"`
	CommentNum  uint32           `json:"comment_num"`
	UserInfo    UserHoleResponse `json:"user_info"`
}
