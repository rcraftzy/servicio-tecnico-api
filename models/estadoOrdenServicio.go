package models

type EstadoOrdenServicio struct {
  ID           int `json:"id" gorm:"primaryKey"`
  State       string `json:"state"`
  EmpresaRefer    int   `json:"empresa_id"`
  Empresa      Empresa `gorm:"foreignKey:EmpresaRefer"`
}
