package models

type UserEmpresa struct {
	ID           int     `json:"id" gorm:"primaryKey"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
	EmpresaRefer int     `json:"empresa_id"`
	Empresa      Empresa `gorm:"foreignKey:EmpresaRefer"`
}
