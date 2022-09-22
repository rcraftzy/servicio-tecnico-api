package models

type EstadoOrdenServicio struct {
  ID           int `json:"id" gorm:"primaryKey"`
  State       string `json:"state"`
  Color       string `json:"color"`
  EmpresaRefer    int   `json:"empresa_id"`
  Empresa      Empresa `gorm:"foreignKey:EmpresaRefer"`
}
