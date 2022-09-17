package models

type Producto struct {
  ID           int `json:"id" gorm:"primaryKey"`
  Codigo       string `json:"codigo"`
  Nombre       string `json:"nombre"`
  PrecioVenta       float64 `json:"precioVenta"`
  StockMin       float64 `json:"stockMin"`
  StockMax       float64 `json:"stockMax"`
  Stock       float64 `json:"stock"`
  ControlaStock bool `json:"controlaStock"`
  AplicaIVA bool `json:"aplicaIva"`
  EmpresaRefer int `json:"empresa_id"`
  Empresa Empresa `gorm:"foreignKey:EmpresaRefer"`
}
