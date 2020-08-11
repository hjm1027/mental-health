package model

//消息物理表
type Message struct {
	Id              uint32 `gorm:"column:id; primary_key" `
	PubUserId       uint32 `gorm:"column:pub_user_id"`
	SubUserId       uint32 `gorm:"column:sub_user_id"`
	Kind            uint8  `gorm:"column:kind"`
	IsRead          bool   `gorm:"column:is_read"`
	Reply           string `gorm:"column:reply"`
	Time            string `gorm:"column:time"`
	HoleId          uint32 `gorm:"column:hole_id"`
	Content         string `gorm:"column:content"`
	Sid             string `gorm:"column:sid"`
	ParentCommentId uint32 `gorm:"column:parent_comment_id"`
}

type MessagePub struct {
	PubUserId       uint32 `json:"pub_user_id"`
	SubUserId       uint32 `json:"sub_user_id"`
	Kind            uint8  `json:"kind"`
	IsRead          bool   `json:"is_read"`
	Reply           string `json:"reply"`
	Time            string `json:"time"`
	HoleId          uint32 `json:"hole_id"`
	Content         string `json:"content"`
	Sid             string `json:"sid"`
	ParentCommentId uint32 `json:"parent_comment_id"`
}

type MessageSub struct {
	UserInfo        UserHoleResponse2 `json:"user_info"`
	Kind            uint8             `json:"kind"`
	IsRead          bool              `json:"is_read"`
	Reply           string            `json:"reply"`
	Time            string            `json:"time"`
	HoleId          uint32            `json:"hole_id"`
	Content         string            `json:"content"`
	Sid             string            `json:"sid"`
	ParentCommentId uint32            `json:"parent_comment_id"`
}
