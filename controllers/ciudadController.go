package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type Ciudad struct {
	ID        int       `json:"id"`
	Nombre    string    `json:"nombre"`
	Provincia Provincia `json:"provincia"`
}

func CreateResponseCiudad(ciudad models.Ciudad, provincia Provincia) Ciudad {
	return Ciudad{ID: ciudad.ID, Nombre: ciudad.Nombre, Provincia: provincia}
}

func CreateCiudad(c *fiber.Ctx) error {
	var ciudad models.Ciudad

	if err := c.BodyParser(&ciudad); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var provincia models.Provincia
	if err := findProvincia(ciudad.ProvinciaRefer, &provincia); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&ciudad)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)

	return c.Status(200).JSON(responseCiudad)
}

func GetCiudades(c *fiber.Ctx) error {
	ciudades := []models.Ciudad{}
	database.DB.Find(&ciudades)
	responseCiudades := []Ciudad{}

	for _, ciudad := range ciudades {
		var provincia models.Provincia
		database.DB.Find(&provincia, "id = ?", ciudad.ProvinciaRefer)
		responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
		responseCiudades = append(responseCiudades, responseCiudad)
	}
	return c.Status(200).JSON(responseCiudades)
}

func FindCiudad(id int, ciudad *models.Ciudad) error {
	database.DB.Find(&ciudad, "id = ?", id)
	if ciudad.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}

func GetCiudad(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var ciudad models.Ciudad

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindCiudad(id, &ciudad); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var provincia models.Provincia

	database.DB.First(&provincia, ciudad.ProvinciaRefer)
	responseProvincia := CreateResponseProvincia(provincia)

	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)

	return c.Status(200).JSON(responseCiudad)
}

func UpdateCiudad(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var ciudad models.Ciudad

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindCiudad(id, &ciudad); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateCiudad struct {
		Nombre         string `json:"nombre"`
		ProvinciaRefer int    `json:"provincia_id"`
	}

	var updateData UpdateCiudad

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	ciudad.Nombre = updateData.Nombre
	ciudad.ProvinciaRefer = updateData.ProvinciaRefer

	var provincia models.Provincia
	if err := findProvincia(ciudad.ProvinciaRefer, &provincia); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Save(&ciudad)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	return c.Status(200).JSON(responseCiudad)
}

func DeleteCiudad(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var ciudad models.Ciudad

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindCiudad(id, &ciudad); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&ciudad).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Delteted product")
}
