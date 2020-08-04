package model

// LoginModel represents a json for registering
type LoginModel struct {
	Sid      string `json:"sid"      binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserModel represents a registered user.
type UserModel struct {
	Id           uint32 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Sid          string `gorm:"column:sid"`
	Username     string `gorm:"column:username"`
	Avatar       string `gorm:"column:avatar"`
	Introduction string `gorm:"column:introduction"`
	Phone        string `gorm:"column:phone"`
	Back_avatar  string `gorm:"column:back_avatar"`
}

// AuthResponse represents a JSON web token.
type AuthResponse struct {
	Token string `json:"token"`
	IsNew uint8  `json:"is_new"`
}

type UserInfoRequest struct {
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Phone        string `json:"phone"`
	Back_avatar  string `json:"back_avatar"`
}

type UserInfoResponse struct {
	Sid          string `json:"sid"`
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Phone        string `json:"phone"`
	Back_avatar  string `json:"back_avatar"`
}

type UserHoleResponse struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
