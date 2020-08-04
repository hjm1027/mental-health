package model

func (hole *HoleModel) TableName() string {
	return "hole"
}

func (hole *HoleLikeModel) TableName() string {
	return "hole_like"
}

func (hole *HoleFavoriteModel) TableName() string {
	return "hole_favorite"
}

func (hole *HoleReadModel) TableName() string {
	return "hole_read"
}

func (hole *HoleModel) New() error {
	d := DB.Self.Create(hole)
	return d.Error
}

func (hole *HoleModel) GetById() error {
	d := DB.Self.Where("id = ?", hole.Id).First(hole)
	return d.Error
}

func (hole *HoleModel) HasLiked(userId uint32) (uint32, bool) {
	var data HoleLikeModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, hole.Id).First(&data)
	return data.Id, !d.RecordNotFound()
}

func (hole *HoleModel) HasFavorited(userId uint32) (uint32, bool) {
	var data HoleFavoriteModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, hole.Id).First(&data)
	return data.Id, !d.RecordNotFound()
}

func (hole *HoleModel) HasRead(userId uint32) (uint32, bool) {
	var data HoleReadModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, hole.Id).First(&data)
	return data.Id, !d.RecordNotFound()
}

func (hole *HoleReadModel) Read() error {
	d := DB.Self.Create(hole)
	data := HoleModel{Id: hole.HoleId}

	err := data.GetById()
	if err != nil {
		return d.Error
	}
	//fmt.Println(data)
	data.ReadNum += 1
	d = DB.Self.Save(&data)
	return d.Error
}
