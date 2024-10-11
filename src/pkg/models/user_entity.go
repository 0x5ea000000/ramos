package models

type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
