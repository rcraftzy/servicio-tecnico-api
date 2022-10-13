package models

type Empresa struct {
	ID            int     `json:"id" gorm:"primaryKey"`
	Ruc           string  `json:"ruc"`
	Nombre        string  `json:"nombre"`
	Direccion     string  `json:"direccion"`
	CiudadRefer   int     `json:"ciudad_id"`
	Ciudad        Ciudad  `gorm:"foreignKey:CiudadRefer"`
	Telefono      string  `json:"telefono"`
	Email         string  `json:"email"`
	PorcentajeIVA float64 `json:"porcentajeIVA"`
}
