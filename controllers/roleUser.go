package controllers 

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type RoleUser struct {
  ID           int `json:"id"`
  Nombre       string `json:"nombre"`
}

func CreateResponseRoleUser(roleUserModel models.RoleUser) RoleUser {
	return RoleUser{ID: roleUserModel.ID, Nombre: roleUserModel.Nombre}
}

func CreateRoleUser(c *fiber.Ctx) error {
	var roleUser models.RoleUser

	if err := c.BodyParser(&roleUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&roleUser)

	responseRoleUser := CreateResponseRoleUser(roleUser)

	return c.Status(200).JSON(responseRoleUser)
}

func GetRolesUser(c *fiber.Ctx) error {
	rolesUser := []models.RoleUser{}
	database.DB.Find(&rolesUser)
	responseRolesUser := []RoleUser{}
	for _, roleUser := range rolesUser {
		responseRoleUser := CreateResponseRoleUser(roleUser)
		responseRolesUser = append(responseRolesUser, responseRoleUser)
	}

	return c.Status(200).JSON(responseRolesUser)
}

func FindRoleUser(id int, roleUser *models.RoleUser) error {
	database.DB.Find(&roleUser, "id = ?", id)
	if roleUser.ID == 0 {
		return errors.New("Role user does not exist")
	}
	return nil
}

func GetRoleUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var roleUser models.RoleUser

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindRoleUser(id, &roleUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseRoleUser := CreateResponseRoleUser(roleUser)

	return c.Status(200).JSON(responseRoleUser)
}

func UpdateRoleUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var roleUser models.RoleUser

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindRoleUser(id, &roleUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateRoleUser struct {
		Nombre         string `json:"nombre"`
	}

	var updateData UpdateRoleUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	roleUser.Nombre = updateData.Nombre

	database.DB.Save(&roleUser)

	responseRoleUser := CreateResponseProvincia(models.Provincia(roleUser))
	return c.Status(200).JSON(responseRoleUser)
}
