package model

type MoodModel struct {
	Id     uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserId uint32 `gorm:"column:user_id"`
	Year   uint32 `gorm:"column:year"`
	Month  uint8  `gorm:"column:month"`
	Day    uint8  `gorm:"column:day"`
	Score  uint8  `gorm:"column:score"`
	Note   string `gorm:"column:note"`
}

type MoodScoreItem struct {
	Day   uint8 `json:"day"`
	Score uint8 `json:"score"`
}
