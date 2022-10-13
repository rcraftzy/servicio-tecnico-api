package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/servicio-tecnico-api/database"
	"github.com/roberto-carlos-tg/servicio-tecnico-api/models"
)

type Provincia struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func CreateResponseProvincia(provinciaModel models.Provincia) Provincia {
	return Provincia{ID: provinciaModel.ID, Nombre: provinciaModel.Nombre}
}

func CreateProvincia(c *fiber.Ctx) error {
	var provincia models.Provincia

	if err := c.BodyParser(&provincia); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&provincia)
	responseProvincia := CreateResponseProvincia(provincia)

	return c.Status(200).JSON(responseProvincia)
}

func GetProvincias(c *fiber.Ctx) error {
	provincias := []models.Provincia{}
	database.DB.Find(&provincias)
	responseProvincias := []Provincia{}
	for _, provincia := range provincias {
		responseProvincia := CreateResponseProvincia(provincia)
		responseProvincias = append(responseProvincias, responseProvincia)
	}

	return c.Status(200).JSON(responseProvincias)
}

func findProvincia(id int, provincia *models.Provincia) error {
	database.DB.Find(&provincia, "id = ?", id)
	if provincia.ID == 0 {
		return errors.New("Product does not exist")
	}
	return nil
}

func GetProvincia(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var provincia models.Provincia

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findProvincia(id, &provincia); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseProvincia(provincia)

	return c.Status(200).JSON(responseUser)
}

func UpdateProvincia(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var provincia models.Provincia

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findProvincia(id, &provincia); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProvincia struct {
		Nombre string `json:"nombre"`
	}

	var updateData UpdateProvincia

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	provincia.Nombre = updateData.Nombre

	database.DB.Save(&provincia)

	responseProduct := CreateResponseProvincia(provincia)
	return c.Status(200).JSON(responseProduct)
}

func DeleteProvincia(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var provincia models.Provincia

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findProvincia(id, &provincia); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&provincia).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Delteted product")
}
