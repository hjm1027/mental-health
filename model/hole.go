package model

func (hole *HoleModel) TableName() string {
	return "hole"
}

func (hole *HoleModel) New() error {
	d := DB.Self.Create(hole)
	return d.Error
}
