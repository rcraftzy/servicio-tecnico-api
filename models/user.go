package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"  gorm:"unique"`
	Email    string `json:"email"`
	RoleUserRefer    int `json:"role_user_id"`
	RoleUser    RoleUser `gorm:"foreignKey:RoleUserRefer"`
	Password []byte `json:"-"`
}
