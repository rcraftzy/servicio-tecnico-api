package models

type Cliente struct {
  ID           int `json:"id" gorm:"primaryKey"`
  Nombres       string `json:"nombres"`
  Apellidos       string `json:"apellidos"`
  Dni       string `json:"dni"`
  Direccion    string   `json:"direccion"`
  Telefono    string   `json:"telefono"`
  Celular    string   `json:"celular"`
  Email    string   `json:"email"`
  Estado    bool   `json:"estado"`
}
