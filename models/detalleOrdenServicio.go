package models

type DetalleOrdenServicio struct {
  ID           int `json:"id" gorm:"primaryKey"`
  OrdenServicioRefer int `json:"orden_servicio_id"`
  OrdenServicio OrdenServicio `gorm:"foreignKey:OrdenServicioRefer"`
  Cantidad float64 `json:"cantidad"`
  ProductoRefer       int `json:"producto_id"`
  Producto Producto `gorm:"foreignKey:ProductoRefer"`
  Descripcion       string `json:"descripcion"`
  PrecioUnitario       float64 `json:"precio_unitario"`
  Descuento    float64   `json:"descuento"`
  PorcentajeIVA    float64   `json:"porcentaje_IVA"`
  ValorIVA    float64   `json:"valor_IVA"`
  Total    float64   `json:"total"`
  EstadoOrdenServicioRefer    int   `json:"estado_orden_servicio_id"`
  EstadoOrdenServicio EstadoOrdenServicio `gorm:"foreignKey:EstadoOrdenServicioRefer"`
}
