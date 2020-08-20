package model

func (u *PosterModel) TableName() string {
	return "poster"
}

// 添加一个海报
func (poster *PosterModel) New() error {
	d := DB.Self.Create(poster)
	return d.Error
}

// 查询最新的海报
func LastestPoster() (*PosterModel, error) {
	var data PosterModel
	d := DB.Self.Table("poster").Order("id DESC").First(&data)
	return &data, d.Error
}
