package controllers 

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type Cliente struct {
  ID           int `json:"id" gorm:"primaryKey"`
  Nombres       string `json:"nombres"`
  Apellidos       string `json:"apellidos"`
  Dni       string `json:"dni"`
  Direccion    string   `json:"direccion"`
  Telefono    string   `json:"telefono"`
  Celular    string   `json:"celular"`
  Email    string   `json:"email"`
  Estado    bool   `json:"estado"`
}

func CreateResponseCliente(clienteModel models.Cliente) Cliente {
  return Cliente{ID: clienteModel.ID, Nombres: clienteModel.Nombres, Apellidos: clienteModel.Apellidos, Dni: clienteModel.Dni, Direccion: clienteModel.Direccion, Telefono: clienteModel.Telefono, Celular: clienteModel.Celular, Email: clienteModel.Email, Estado: clienteModel.Estado}
}

func CreateCliente(c *fiber.Ctx) error {
	var cliente models.Cliente

	if err := c.BodyParser(&cliente); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&cliente)
	responseCliente := CreateResponseCliente(cliente)

	return c.Status(200).JSON(responseCliente)
}

func GetClientes(c *fiber.Ctx) error {
	clientes := []models.Cliente{}
	database.DB.Find(&clientes)
	responseClientes := []Cliente{}
	for _, cliente := range clientes {
		responseCliente := CreateResponseCliente(cliente)
		responseClientes = append(responseClientes, responseCliente)
	}

	return c.Status(200).JSON(responseClientes)
}

func FindCliente(id int, cliente *models.Cliente) error {
	database.DB.Find(&cliente, "id = ?", id)
	if cliente.ID == 0 {
		return errors.New("Product does not exist")
	}
	return nil
}

func GetCliente(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var cliente models.Cliente

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindCliente(id, &cliente); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseCliente(cliente)

	return c.Status(200).JSON(responseUser)
}

func UpdateCliente(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var cliente models.Cliente

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindCliente(id, &cliente); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateCliente struct {
		Nombres         string `json:"nombre"`
    Apellidos       string `json:"apellidos"`
    Dni       string `json:"dni"`
    Direccion    string   `json:"direccion"`
    Telefono    string   `json:"telefono"`
    Celular    string   `json:"celular"`
    Email    string   `json:"email"`
    Estado    bool   `json:"estado"`
	}

	var updateData UpdateCliente

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	cliente.Nombres = updateData.Nombres
	cliente.Apellidos = updateData.Apellidos
	cliente.Dni = updateData.Dni
	cliente.Direccion = updateData.Direccion
	cliente.Telefono = updateData.Telefono
	cliente.Celular = updateData.Celular
	cliente.Email = updateData.Email
	cliente.Estado = updateData.Estado

	database.DB.Save(&cliente)

	responseCliente := CreateResponseCliente(cliente)
	return c.Status(200).JSON(responseCliente)
}
