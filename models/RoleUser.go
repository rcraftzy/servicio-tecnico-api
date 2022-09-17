package models

type RoleUser struct {
  ID           int `json:"id" gorm:"primaryKey"`
	Nombre     string `json:"nombre"  gorm:"unique"`
}
