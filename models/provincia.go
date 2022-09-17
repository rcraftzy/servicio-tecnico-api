package models

type Provincia struct {
  ID           int `json:"id" gorm:"primaryKey"`
  Nombre       string `json:"nombre"`
}
