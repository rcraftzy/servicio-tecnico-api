package controllers 

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type EstadoOrdenServicio struct {
  ID           int `json:"id" gorm:"primaryKey"`
  State       string `json:"state"`
  Color       string `json:"color"`
  Empresa      Empresa `json:"empresa"`
}

func CreateResponseEstadoOrdenServicio(estadoOrdenServicioModel models.EstadoOrdenServicio, empresa Empresa) EstadoOrdenServicio {
  return EstadoOrdenServicio {ID: estadoOrdenServicioModel.ID ,State: estadoOrdenServicioModel.State ,Color: estadoOrdenServicioModel.Color, Empresa: empresa}
}

func CreateEstadoOrdenServicio(c *fiber.Ctx) error {
	var estadoOrdenServicio models.EstadoOrdenServicio

	if err := c.BodyParser(&estadoOrdenServicio); err != nil {
		return c.Status(400).JSON(err.Error())
	}

  var empresa models.Empresa
	if err := findEmpresa(estadoOrdenServicio.EmpresaRefer, &empresa); err != nil {
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

	database.DB.Create(&estadoOrdenServicio)

  responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
  responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
  responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estadoOrdenServicio, responseEmpresa)

	return c.Status(200).JSON(responseEstadoOrdenServicio)
}

func GetEstadosOrdenServicio(c *fiber.Ctx) error {
	estadosOrdenServicio := []models.EstadoOrdenServicio{}
	database.DB.Find(&estadosOrdenServicio)
	responseEstadosOrdenServicio := []EstadoOrdenServicio{}

	for _, estadoOrdenServicio := range estadosOrdenServicio {

    var empresa models.Empresa
		database.DB.Find(&empresa, "id = ?", estadoOrdenServicio.EmpresaRefer)

    var ciudad models.Ciudad
		database.DB.Find(&ciudad, "id = ?", empresa.CiudadRefer)

		var provincia models.Provincia
		database.DB.Find(&provincia, "id = ?", ciudad.ProvinciaRefer)

    responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
		responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
    responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estadoOrdenServicio, responseEmpresa)
    responseEstadosOrdenServicio = append(responseEstadosOrdenServicio, responseEstadoOrdenServicio)
	}
	return c.Status(200).JSON(responseEstadosOrdenServicio)
}

func FindEstadoOrdenServicio(id int, estadoOrdenServicio *models.EstadoOrdenServicio) error {
	database.DB.Find(&estadoOrdenServicio, "id = ?", id)
	if estadoOrdenServicio.ID == 0 {
		return errors.New("Estado does not exist")
	}
	return nil
}

func GetEstadoOrdenServicio(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var estadoOrdenServicio models.EstadoOrdenServicio

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindEstadoOrdenServicio(id, &estadoOrdenServicio); err != nil {
		return c.Status(400).JSON(err.Error())
	}

  var empresa models.Empresa
	database.DB.First(&empresa, estadoOrdenServicio.EmpresaRefer)
  
	var ciudad models.Ciudad
	database.DB.First(&ciudad, empresa.CiudadRefer)

	var provincia models.Provincia
	database.DB.First(&provincia, ciudad.ProvinciaRefer)

  responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
  responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estadoOrdenServicio, responseEmpresa)

	return c.Status(200).JSON(responseEstadoOrdenServicio)
}

func FindEstadoOrdenServicioByEmpresa(empresa_id int, estadoOrdenServicio *models.EstadoOrdenServicio) error {
  database.DB.Find(&estadoOrdenServicio, "empresa_refer = ?", empresa_id)
	if estadoOrdenServicio.EmpresaRefer == 0 {
		return errors.New("Empresa does not exist")
	}
	return nil
}

func GetEstadoOrdenServicioByEmpresa(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

  estadosOrdenServicio := []models.EstadoOrdenServicio{}
	database.DB.Find(&estadosOrdenServicio)
	responseEstadosOrdenServicio := []EstadoOrdenServicio{}

	for _, estadoOrdenServicio := range estadosOrdenServicio {

    if err != nil {
      return c.Status(400).JSON("Please ensure that :id is an integer")
    }

    if err := FindEstadoOrdenServicioByEmpresa(id, &estadoOrdenServicio); err != nil {
      return c.Status(400).JSON(err.Error())
    }

    var empresa models.Empresa
		database.DB.Find(&empresa, "id = ?", estadoOrdenServicio.EmpresaRefer)

    var ciudad models.Ciudad
		database.DB.Find(&ciudad, "id = ?", empresa.CiudadRefer)

		var provincia models.Provincia
		database.DB.Find(&provincia, "id = ?", ciudad.ProvinciaRefer)

    responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
		responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
    responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estadoOrdenServicio, responseEmpresa)
    responseEstadosOrdenServicio = append(responseEstadosOrdenServicio, responseEstadoOrdenServicio)
	}
 	return c.Status(200).JSON(responseEstadosOrdenServicio)
}

func UpdateEstadoOrdenServicio(c *fiber.Ctx) error {
  id, err := c.ParamsInt("id")

	var estado models.EstadoOrdenServicio

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

  if err := FindEstadoOrdenServicio(id, &estado); err != nil {
		return c.Status(400).JSON(err.Error())
	}

  type UpdateEmpresa struct {
    State           string `json:"state"`
    Color           string `json:"color"`
    EmpresaRefer    int `json:"empresa_id"`
  }

	var updateData UpdateEmpresa

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

  estado.State = updateData.State
  estado.Color = updateData.Color
  estado.EmpresaRefer = updateData.EmpresaRefer

	var empresa models.Empresa
	if err := findEmpresa(estado.EmpresaRefer, &empresa); err != nil {
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

	database.DB.Save(&estado)

  responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estado, responseEmpresa)

	return c.Status(200).JSON(responseEstadoOrdenServicio)
}

func DeleteEstadoOrdenServicio(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var estado models.EstadoOrdenServicio

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindEstadoOrdenServicio(id, &estado); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&estado).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Delteted product")
}
