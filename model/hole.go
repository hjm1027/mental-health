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

// 判断问题是否已经被当前用户点赞
func HasLiked(userId uint32, holeId uint32) (uint32, bool) {
	var data HoleLikeModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, holeId).First(&data)
	return data.Id, !d.RecordNotFound()
}

// 点赞问题
func Like(userId uint32, holeId uint32) error {
	var data = HoleLikeModel{
		HoleId: holeId,
		UserId: userId,
	}
	d := DB.Self.Create(&data)
	return d.Error
}

// 取消点赞
func Unlike(id uint32) error {
	var data = HoleLikeModel{Id: id}
	d := DB.Self.Delete(&data)
	return d.Error
}

// 判断问题是否已经被当前用户收藏
func HasFavorited(userId uint32, holeId uint32) (uint32, bool) {
	var data HoleFavoriteModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, holeId).First(&data)
	return data.Id, !d.RecordNotFound()
}

// 收藏问题
func Favorite(userId uint32, holeId uint32) error {
	var data = HoleFavoriteModel{
		HoleId: holeId,
		UserId: userId,
	}
	d := DB.Self.Create(&data)
	return d.Error
}

// 取消收藏
func Unfavorite(id uint32) error {
	var data = HoleFavoriteModel{Id: id}
	d := DB.Self.Delete(&data)
	return d.Error
}
