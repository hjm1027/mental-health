package model

// 必须要写一个TableName函数，返回table的名字，否则gorm读取不到表。
func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserCodeModel) TableName() string {
	return "user_code"
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

// 检验是否存在用户openid
func HaveUserCode(userId uint32) (uint8, error) {
	d := DB.Self.Table("user_code").Where("user_id = ?", userId)
	if d.RecordNotFound() {
		return 0, nil
	}
	return 1, d.Error
}

// 通过id获取用户信息
func (u *UserModel) GetUserById() error {
	d := DB.Self.Where("id = ?", u.Id).First(&u)
	return d.Error
}

// 通过学号获取用户信息
func (u *UserModel) GetUserBySid() error {
	d := DB.Self.Where("sid = ?", u.Sid).First(&u)
	return d.Error
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

// 通过用户id获取用户信息
func GetUserInfoById(id uint32) (*UserInfoResponse, error) {
	u, err := GetUserById(id)
	if err != nil {
		return &UserInfoResponse{}, err
	}
	info := u.GetInfo()
	return info, nil
}

func (u *UserModel) GetInfo() *UserInfoResponse {
	info := UserInfoResponse{
		Sid:          u.Sid,
		Username:     u.Username,
		Avatar:       u.Avatar,
		Introduction: u.Introduction,
		Phone:        u.Phone,
		Back_avatar:  u.Back_avatar,
	}
	return &info
}

func (user *UserCodeModel) Create() error {
	d := DB.Self.Create(user)
	return d.Error
}

func (user *UserCodeModel) Save() error {
	var data UserCodeModel
	d := DB.Self.Table("user_code").Where("user_id = ?", user.UserId).First(&data)
	data.Openid = user.Openid
	data.Unionid = user.Unionid
	d = DB.Self.Save(data)
	return d.Error
}
