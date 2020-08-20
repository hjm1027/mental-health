package model

type PosterModel struct {
	Id       uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Home     string `gorm:"column:home"`
	Platform string `gorm:"column:platform"`
	Hole     string `gorm:"column:hole"`
}

type PosterResponse struct {
	Home     string `json:"home"`
	Platform string `json:"platform"`
	Hole     string `json:"hole"`
}
