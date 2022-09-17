package models

type Ciudad struct {
  ID           int `json:"id" gorm:"primaryKey"`
  Nombre       string `json:"nombre"`
  ProvinciaRefer    int   `json:"provincia_id"`
  Provincia      Provincia `gorm:"foreignKey:ProvinciaRefer"`
}
