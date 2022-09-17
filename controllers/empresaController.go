package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type Empresa struct {
  ID           int `json:"id"`
  Ruc          string `json:"ruc"`
  Nombre       string `json:"nombre"`
  Direccion    string `json:"direccion"`
  Ciudad      Ciudad `json:"ciudad"`
  Telefono   string `json:"telefono"`
  Email   string `json:"email"`
  PorcentajeIVA   float64 `json:"porcentajeIVA"`
}

func CreateResponseEmpresa(empresa models.Empresa, ciudad Ciudad) Empresa {
  return Empresa {ID: empresa.ID, Ruc: empresa.Ruc, Nombre: empresa.Nombre, Direccion: empresa.Direccion, Ciudad: ciudad, Telefono: empresa.Telefono, Email: empresa.Email, PorcentajeIVA: empresa.PorcentajeIVA}
}

func CreateEmpresa(c *fiber.Ctx) error {
	var empresa models.Empresa

	if err := c.BodyParser(&empresa); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var ciudad models.Ciudad
	if err := FindCiudad(empresa.CiudadRefer, &ciudad); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var provincia models.Provincia
	if err := findProvincia(ciudad.ProvinciaRefer, &provincia); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&empresa)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)

	return c.Status(200).JSON(responseEmpresa)
}

func GetEmpresas(c *fiber.Ctx) error {
	empresas := []models.Empresa{}
	database.DB.Find(&empresas)
	responseEmpresas := []Empresa{}

	for _, empresa := range empresas {
		var ciudad models.Ciudad
		database.DB.Find(&ciudad, "id = ?", empresa.CiudadRefer)

    var provincia models.Provincia
		database.DB.Find(&provincia, "id = ?", ciudad.ProvinciaRefer)

		responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
		responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
		responseEmpresas = append(responseEmpresas, responseEmpresa)
	}
	return c.Status(200).JSON(responseEmpresas)
}

func findEmpresa(id int, empresa *models.Empresa) error {
	database.DB.Find(&empresa, "id = ?", id)
	if empresa.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}

func GetEmpresa(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var empresa models.Empresa

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findEmpresa(id, &empresa); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var ciudad models.Ciudad
	database.DB.First(&ciudad, empresa.CiudadRefer)

	var provincia models.Provincia
	database.DB.First(&empresa, ciudad.ProvinciaRefer)

  responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)

	return c.Status(200).JSON(responseEmpresa)
}
