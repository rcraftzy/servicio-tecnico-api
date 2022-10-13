package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type Role struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func CreateResponseRole(roleModel models.Role) Role {
	return Role{ID: roleModel.ID, Nombre: roleModel.Nombre}
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&role)

	responseRoleUser := CreateResponseRole(role)

	return c.Status(200).JSON(responseRoleUser)
}

func GetRoles(c *fiber.Ctx) error {
	roles := []models.Role{}
	database.DB.Find(&roles)
	responseRoles := []Role{}
	for _, role := range roles {
		responseRole := CreateResponseRole(role)
		responseRoles = append(responseRoles, responseRole)
	}

	return c.Status(200).JSON(responseRoles)
}

func FindRole(id int, role *models.Role) error {
	database.DB.Find(&role, "id = ?", id)
	if role.ID == 0 {
		return errors.New("Role user does not exist")
	}
	return nil
}

func GetRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var role models.Role

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindRole(id, &role); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseRole := CreateResponseRole(role)

	return c.Status(200).JSON(responseRole)
}

func UpdateRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var role models.Role

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindRole(id, &role); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateRole struct {
		Nombre string `json:"nombre"`
	}

	var updateData UpdateRole

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	role.Nombre = updateData.Nombre

	database.DB.Save(&role)

	responseRole := CreateResponseRole(role)
	return c.Status(200).JSON(responseRole)
}
