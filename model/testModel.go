package model

type TestModel struct {
	Id      uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Url     string `gorm:"column:url"`
	Header  string `gorm:"column:header"`
	Content string `gorm:"column:content"`
	Image   string `gorm:"column:image"`
}
