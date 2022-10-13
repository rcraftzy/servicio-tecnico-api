package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/servicio-tecnico-api/database"
	"github.com/roberto-carlos-tg/servicio-tecnico-api/models"
)

type RoleUser struct {
	ID   int          `json:"id"`
	User UserResponse `json:"user"`
	Role Role         `json:"role"`
}

func CreateResponseRoleUser(roleUserModel models.RoleUser, user UserResponse, role Role) RoleUser {
	return RoleUser{ID: roleUserModel.ID, User: user, Role: role}
}

func CreateRoleUser(c *fiber.Ctx) error {
	var roleUser models.RoleUser

	if err := c.BodyParser(&roleUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := FindUser(roleUser.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var role models.Role
	if err := FindRole(roleUser.RoleRefer, &role); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&roleUser)

	responseUser := CreateResponseUser(user)
	responseRole := CreateResponseRole(role)
	responseRoleUser := CreateResponseRoleUser(roleUser, responseUser, responseRole)

	return c.Status(200).JSON(responseRoleUser)
}

func GetRolesUser(c *fiber.Ctx) error {
	rolesUser := []models.RoleUser{}
	database.DB.Find(&rolesUser)
	responseRolesUser := []RoleUser{}
	for _, roleUser := range rolesUser {

		var user models.User
		database.DB.Find(&user, "id = ?", roleUser.UserRefer)

		var role models.Role
		database.DB.Find(&role, "id = ?", roleUser.RoleRefer)

		responseUser := CreateResponseUser(user)
		responseRole := CreateResponseRole(role)
		responseRoleUser := CreateResponseRoleUser(roleUser, responseUser, responseRole)
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
	var user models.User
	var role models.Role

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindRoleUser(id, &roleUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	responseRole := CreateResponseRole(role)
	responseRoleUser := CreateResponseRoleUser(roleUser, responseUser, responseRole)

	return c.Status(200).JSON(responseRoleUser)
}

func UpdateRoleUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var roleUser models.RoleUser
	var user models.User
	var role models.Role

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindRoleUser(id, &roleUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateRoleUser struct {
		UserRefer int `json:"user_id"`
		RoleRefer int `json:"role_id"`
	}

	var updateData UpdateRoleUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	roleUser.UserRefer = updateData.UserRefer
	roleUser.RoleRefer = updateData.RoleRefer

	database.DB.Save(&roleUser)

	responseUser := CreateResponseUser(user)
	responseRole := CreateResponseRole(role)
	responseRoleUser := CreateResponseRoleUser(roleUser, responseUser, responseRole)
	return c.Status(200).JSON(responseRoleUser)
}
