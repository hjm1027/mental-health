package model

import "time"

type ReserveModel struct {
	Id          uint32    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Weekday     uint8     `gorm:"column:weekday"`
	Schedule    uint8     `gorm:"column:schedule"`
	Teacher     string    `gorm:"column:teacher"`
	Reserve     uint8     `gorm:"column:reserve"`
	Time        time.Time `gorm:"column:time"`
	AdvanceTime uint8     `gorm:"column:advance_time"`
	Type        uint8     `gorm:"column:type"`
	Method      uint8     `gorm:"column:method"`
	UserId      uint32    `gorm:"column:user_id"`
}
