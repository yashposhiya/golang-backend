package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
