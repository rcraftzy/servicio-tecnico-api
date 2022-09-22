package models

type RoleUser struct {
  ID           int `json:"id" gorm:"primaryKey"`
  UserRefer    int   `json:"user_id"`
  User        User      `gorm:"foreignKey:UserRefer"`
  RoleRefer    int   `json:"role_id"`
  Role  Role       `gorm:"foreignKey:RoleRefer"`
}
