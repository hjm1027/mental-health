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
