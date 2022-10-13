package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserEmpresa struct {
	ID      int          `json:"id"`
	User    UserResponse `json:"user"`
	Empresa Empresa      `json:"empresa"`
}

func CreateResponseUser(user models.User) UserResponse {
	return UserResponse{Id: user.Id, Name: user.Name, Email: user.Email}
}

func CreateResponseUserEmpresa(userEmpresa models.UserEmpresa, user UserResponse, empresa Empresa) UserEmpresa {
	return UserEmpresa{ID: userEmpresa.ID, User: user, Empresa: empresa}
}

func CreateUserEmpresa(c *fiber.Ctx) error {
	var userEmpresa models.UserEmpresa

	if err := c.BodyParser(&userEmpresa); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := FindUser(userEmpresa.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var empresa models.Empresa
	if err := findEmpresa(userEmpresa.EmpresaRefer, &empresa); err != nil {
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

	database.DB.Create(&userEmpresa)

	responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseUser := CreateResponseUser(user)
	responseUserEmpresa := CreateResponseUserEmpresa(userEmpresa, responseUser, responseEmpresa)

	return c.Status(200).JSON(responseUserEmpresa)
}

func GetUsersEmpresas(c *fiber.Ctx) error {
	usersEmpresas := []models.UserEmpresa{}
	database.DB.Find(&usersEmpresas)
	responseUsersEmpresas := []UserEmpresa{}

	for _, userEmpresa := range usersEmpresas {

		var user models.User
		database.DB.Find(&user, "id = ?", userEmpresa.UserRefer)

		var empresa models.Empresa
		database.DB.Find(&empresa, "id = ?", userEmpresa.EmpresaRefer)

		var ciudad models.Ciudad
		database.DB.Find(&ciudad, "id = ?", empresa.CiudadRefer)

		var provincia models.Provincia
		database.DB.Find(&provincia, "id = ?", ciudad.ProvinciaRefer)

		responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
		responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
		responseUser := CreateResponseUser(user)
		responseUserEmpresa := CreateResponseUserEmpresa(userEmpresa, responseUser, responseEmpresa)
		responseUsersEmpresas = append(responseUsersEmpresas, responseUserEmpresa)
	}
	return c.Status(200).JSON(responseUsersEmpresas)
}

func FindUserEmpresa(id int, ordenServicio *models.OrdenServicio) error {
	database.DB.Find(&ordenServicio, "id = ?", id)
	if ordenServicio.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}
func FindUserEmpresaUser(id int, userEmpresa *models.UserEmpresa) error {
	database.DB.Find(&userEmpresa, "id = ?", id)
	if userEmpresa.UserRefer == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}
func GetUserEmpresa(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var userEmpresa models.UserEmpresa

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindUserEmpresaUser(id, &userEmpresa); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	database.DB.First(&user, userEmpresa.UserRefer)

	var empresa models.Empresa
	database.DB.First(&empresa, userEmpresa.EmpresaRefer)

	var ciudad models.Ciudad
	database.DB.First(&ciudad, empresa.CiudadRefer)

	var provincia models.Provincia
	database.DB.First(&provincia, ciudad.ProvinciaRefer)

	responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseUser := CreateResponseUser(user)
	responseUserEmpresa := CreateResponseUserEmpresa(userEmpresa, responseUser, responseEmpresa)

	return c.Status(200).JSON(responseUserEmpresa)
}
