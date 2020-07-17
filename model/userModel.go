package model

// LoginModel represents a json for registering
type LoginModel struct {
	Sid      string `json:"sid"      binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserModel represents a registered user.
type UserModel struct {
	Id       uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Sid      string `gorm:"column:sid"`
	Username string `gorm:"column:username"`
	Avatar   string `gorm:"column:avatar"`
}
