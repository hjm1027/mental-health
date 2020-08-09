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

//问题返回格式
type HoleInfoResponse struct {
	HoleId      uint32           `json:"hole_id"`
	Type        uint8            `json:"type"`
	Name        string           `json:"name"`
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

type HoleInfoResponse2 struct {
	HoleId      uint32            `json:"hole_id"`
	Type        uint8             `json:"type"`
	Content     string            `json:"content"`
	LikeNum     uint32            `json:"like_num"`
	ReadNum     uint32            `json:"read_num"`
	FavoriteNum uint32            `json:"favorite_num"`
	IsLike      bool              `json:"is_like"`
	IsFavorite  bool              `json:"is_favorite"`
	Time        time.Time         `json:"time"`
	CommentNum  uint32            `json:"comment_num"`
	UserInfo    UserHoleResponse2 `json:"user_info"`
}

// 父评论物理表
type ParentCommentModel struct {
	Id            uint32    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserId        uint32    `gorm:"column:user_id"`
	HoleId        uint32    `gorm:"column:hole_id"`
	Content       string    `gorm:"column:content"`
	Time          time.Time `gorm:"column:time"`
	SubCommentNum uint32    `gorm:"column:sub_comment_num"`
}

// 子评论物理表
type SubCommentModel struct {
	Id           uint32    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserId       uint32    `gorm:"column:user_id"`
	TargetUserId uint32    `gorm:"column:target_user_id"` // 回复的用户id
	ParentId     uint32    `gorm:"column:parent_id"`
	Content      string    `gorm:"column:content"`
	Time         time.Time `gorm:"column:time"`
}

// 评论点赞中间表
type CommentLikeModel struct {
	Id        uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserId    uint32 `gorm:"column:user_id"`
	CommentId uint32 `gorm:"column:comment_id"`
}

// 评论信息
type CommentInfo struct {
	Id             uint32            `json:"id"`
	Content        string            `json:"content"`
	LikeNum        uint32            `json:"like_num"`
	IsLike         bool              `json:"is_like"`
	Date           string            `json:"date"`
	Time           time.Time         `json:"time"`
	UserInfo       *UserInfoResponse `json:"user_info"`
	TargetUserInfo *UserInfoResponse `json:"target_user_info"`
}

// 返回的评论列表，一级评论模型
type ParentCommentInfo struct {
	Id             uint32            `json:"id"` // 父评论id
	Content        string            `json:"content"`
	LikeNum        uint32            `json:"like_num"`
	IsLike         bool              `json:"is_like"`
	Time           time.Time         `json:"time"`
	UserInfo       UserHoleResponse2 `json:"user_info"`
	SubCommentsNum uint32            `json:"sub_comments_num"`
}

// 返回的评论列表，一级评论模型
type ParentCommentInfo2 struct {
	Id              uint32             `json:"id"` // 父评论id
	Content         string             `json:"content"`
	LikeNum         uint32             `json:"like_num"`
	IsLike          bool               `json:"is_like"`
	Time            time.Time          `json:"time"`
	UserInfo        *UserHoleResponse2 `json:"user_info"`
	SubCommentsNum  uint32             `json:"sub_comments_num"`
	SubCommentsList *[]CommentInfo     `json:"sub_comments_list"`
}

type TargetUserInfo struct {
	Id uint32 `json:"id"`
}
