package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"  gorm:"unique"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}
