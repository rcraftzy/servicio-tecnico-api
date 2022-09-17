package models

type OrdenServicio struct {
  ID           int `json:"id" gorm:"primaryKey"`
  NumOrden       string `json:"numOrden"`
  EmpresaRefer    int   `json:"empresa_id"`
  Empresa      Empresa `gorm:"foreignKey:EmpresaRefer"`
  FechaEmision string `json:"fecha_emision"`
  EstadoOrdenServicioRefer    int   `json:"estado_orden_servicio_id"`
  EstadoOrdenServicio EstadoOrdenServicio `gorm:"foreignKey:EstadoOrdenServicioRefer"`
  SubTotalConIVA    float64   `json:"sub_total_con_IVA"`
  SubTotalSinIVA    float64   `json:"sub_total_sin_IVA"`
  TecnicoRefer int `json:"tecnico_id"`
  Tecnico Tecnico `gorm:"foreignKey:TecnicoRefer"`
  Descuento    float64   `json:"descuento"`
  ValorIVA    float64   `json:"valor_IVA"`
  Total    float64   `json:"total"`
  Observaciones    string   `json:"observaciones"`
}
