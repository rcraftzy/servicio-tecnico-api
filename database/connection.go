package database

import (
	"github.com/roberto-carlos-tg/go-auht/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:roca-01@/goauth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Ciudad{}, &models.Producto{}, &models.OrdenServicio{}, &models.Tecnico{}, &models.Provincia{}, &models.EstadoOrdenServicio{}, &models.Empresa{}, &models.DetalleOrdenServicio{}, &models.Cliente{}, &models.UserEmpresa{}, &models.RoleUser{})
}
