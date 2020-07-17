package model

// 必须要写一个TableName函数，返回table的名字，否则gorm读取不到表。
func (u *UserModel) TableName() string {
	return "user"
}

// Create creates a new user account.
func (user *UserModel) Create() error {
	d := DB.Self.Create(user)
	return d.Error
}

// HaveUser determines whether there is this user or not by the user identifier.
func (user *UserModel) HaveUser() (uint8, error) {
	d := DB.Self.First(user, "sid = ?", user.Sid)
	if d.RecordNotFound() {
		return 0, nil
	}
	return 1, d.Error
}
