package models

type Tecnico struct {
  ID           int `json:"id" gorm:"primaryKey"`
  Cedula       string `json:"cedula"`
  Nombre       string `json:"nombre"`
  Apellido       string `json:"apellido"`
  Email       string `json:"email"`
  Telefono       string `json:"telefono"`
  Direccion       string `json:"direccion"`
  CiudadRefer    int   `json:"ciudad_id"`
  Ciudad      Ciudad `gorm:"foreignKey:CiudadRefer"`
  EmpresaRefer int `json:"empresa_id"`
  Empresa Empresa `gorm:"foreignKey:EmpresaRefer"`
}
