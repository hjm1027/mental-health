package model

// 必须要写一个TableName函数，返回table的名字，否则gorm读取不到表。
func (u *UserModel) TableName() string {
	return "user"
}

// 添加一个用户
func (user *UserModel) Create() error {
	d := DB.Self.Create(user)
	return d.Error
}

// 检验是否存在用户
func (user *UserModel) HaveUser() (uint8, error) {
	d := DB.Self.First(user, "sid = ?", user.Sid)
	if d.RecordNotFound() {
		return 0, nil
	}
	return 1, d.Error
}

// 通过学号获取用户
func GetUserBySid(sid string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("sid = ?", sid).First(&u)
	return u, d.Error
}

// 通过id获取用户
func GetUserById(id uint32) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ?", id).First(&u)
	return u, d.Error
}

//更新用户信息
func (u *UserModel) UpdateInfo(info *UserInfoRequest) error {
	u.Username = info.Username
	u.Avatar = info.Avatar
	u.Introduction = info.Introduction
	u.Phone = info.Phone
	u.Back_avatar = info.Back_avatar
	return DB.Self.Save(u).Error
}

// 通过用户id更新用户信息
func UpdateInfoById(id uint32, info *UserInfoRequest) error {
	u, err := GetUserById(id)
	if err != nil {
		return err
	}
	if err = u.UpdateInfo(info); err != nil {
		return err
	}
	return nil
}
