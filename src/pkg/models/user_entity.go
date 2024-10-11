package models

type User struct {
	BaseModel
	Username string `gorm:"unique"`
	Password string
}
