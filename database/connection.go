package database

import (
	"os"

	"github.com/roberto-carlos-tg/servicio-tecnico-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_USER")+":"+os.Getenv("DATABASE_PASSWORD")+"@/"+os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Ciudad{}, &models.Producto{}, &models.OrdenServicio{}, &models.Tecnico{}, &models.Provincia{}, &models.EstadoOrdenServicio{}, &models.Empresa{}, &models.DetalleOrdenServicio{}, &models.Cliente{}, &models.UserEmpresa{}, &models.RoleUser{})
}
