package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/servicio-tecnico-api/database"
	"github.com/roberto-carlos-tg/servicio-tecnico-api/models"
)

type Tecnico struct {
	ID        int     `json:"id"`
	Cedula    string  `json:"cedula"`
	Nombre    string  `json:"nombre"`
	Apellido  string  `json:"apellido"`
	Email     string  `json:"email"`
	Telefono  string  `json:"telefono"`
	Direccion string  `json:"direccion"`
	Ciudad    Ciudad  `json:"ciudad"`
	Empresa   Empresa `json:"empresa"`
}

func CreateResponseTecnico(tecnico models.Tecnico, ciudad Ciudad, empresa Empresa) Tecnico {
	return Tecnico{ID: tecnico.ID, Cedula: tecnico.Cedula, Nombre: tecnico.Nombre, Apellido: tecnico.Apellido, Email: tecnico.Email, Telefono: tecnico.Telefono, Direccion: tecnico.Direccion, Ciudad: ciudad, Empresa: empresa}
}

func CreateTecnico(c *fiber.Ctx) error {
	var tecnico models.Tecnico

	if err := c.BodyParser(&tecnico); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var ciudad models.Ciudad
	if err := FindCiudad(tecnico.CiudadRefer, &ciudad); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var provincia models.Provincia
	if err := findProvincia(ciudad.ProvinciaRefer, &provincia); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var empresa models.Empresa
	if err := findEmpresa(tecnico.EmpresaRefer, &empresa); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&tecnico)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseTecnico := CreateResponseTecnico(tecnico, responseCiudad, responseEmpresa)

	return c.Status(200).JSON(responseTecnico)
}

func GetTecnicos(c *fiber.Ctx) error {
	tecnicos := []models.Tecnico{}
	database.DB.Find(&tecnicos)
	responseTecnicos := []Tecnico{}

	for _, tecnico := range tecnicos {

		var empresa models.Empresa
		database.DB.Find(&empresa, "id = ?", tecnico.EmpresaRefer)

		var ciudad models.Ciudad
		database.DB.Find(&ciudad, "id = ?", empresa.CiudadRefer)

		var provincia models.Provincia
		database.DB.Find(&provincia, "id = ?", ciudad.ProvinciaRefer)

		responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
		responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
		responseTecnico := CreateResponseTecnico(tecnico, responseCiudad, responseEmpresa)
		responseTecnicos = append(responseTecnicos, responseTecnico)
	}
	return c.Status(200).JSON(responseTecnicos)
}

func FindTecnico(id int, tecnico *models.Tecnico) error {
	database.DB.Find(&tecnico, "id = ?", id)
	if tecnico.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}

func GetTecnico(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var tecnico models.Tecnico

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindTecnico(id, &tecnico); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var empresa models.Empresa
	database.DB.First(&empresa, tecnico.EmpresaRefer)

	var ciudad models.Ciudad
	database.DB.First(&ciudad, empresa.CiudadRefer)

	var provincia models.Provincia
	database.DB.First(&provincia, ciudad.ProvinciaRefer)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseTecnico := CreateResponseTecnico(tecnico, responseCiudad, responseEmpresa)

	return c.Status(200).JSON(responseTecnico)
}

func UpdateTecnico(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var tecnico models.Tecnico

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindTecnico(id, &tecnico); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateTecnico struct {
		Cedula       string `json:"cedula"`
		Nombre       string `json:"nombre"`
		Apellido     string `json:"apellido"`
		Email        string `json:"email"`
		Telefono     string `json:"telefono"`
		Direccion    string `json:"direccion"`
		CiudadRefer  int    `json:"ciudad_id"`
		EmpresaRefer int    `json:"empresa_id"`
	}

	var updateData UpdateTecnico

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	tecnico.Cedula = updateData.Cedula
	tecnico.Nombre = updateData.Nombre
	tecnico.Apellido = updateData.Apellido
	tecnico.Email = updateData.Email
	tecnico.Telefono = updateData.Telefono
	tecnico.Direccion = updateData.Direccion
	tecnico.CiudadRefer = updateData.CiudadRefer
	tecnico.EmpresaRefer = updateData.EmpresaRefer

	database.DB.Save(&tecnico)

	var empresa models.Empresa
	database.DB.First(&empresa, tecnico.EmpresaRefer)

	var ciudad models.Ciudad
	database.DB.First(&ciudad, empresa.CiudadRefer)

	var provincia models.Provincia
	database.DB.First(&provincia, ciudad.ProvinciaRefer)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseCliente := CreateResponseTecnico(tecnico, responseCiudad, responseEmpresa)
	return c.Status(200).JSON(responseCliente)
}
