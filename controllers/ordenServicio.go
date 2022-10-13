package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type OrdenServicio struct {
	ID                   int                 `json:"id"`
	NumOrden             string              `json:"numOrden"`
	Empresa              Empresa             `json:"empresa"`
	FechaEmision         string              `json:"fecha_emision"`
	EstadoOrdenServicio  EstadoOrdenServicio `json:"estado_orden_servicio"`
	SubTotalConIVA       float64             `json:"sub_total_con_IVA"`
	SubTotalSinIVA       float64             `json:"sub_total_sin_IVA"`
	Tecnico              Tecnico             `json:"tecnico"`
	Cliente              Cliente             `json:"cliente"`
	Descuento            float64             `json:"descuento"`
	ValorIVA             float64             `json:"valor_IVA"`
	Total                float64             `json:"total"`
	Observaciones        string              `json:"observaciones"`
	DiagnosticoRecepcion string              `json:"diagnostico_recepcion"`
	DiagnosticoTecnico   string              `json:"diagnostico_tecnico"`
}

func CreateResponseOrdenServicio(ordenServicioModel models.OrdenServicio, empresa Empresa, estadoOrdenServicio EstadoOrdenServicio, tecnico Tecnico, cliente Cliente) OrdenServicio {
	return OrdenServicio{ID: ordenServicioModel.ID, NumOrden: ordenServicioModel.NumOrden, Empresa: empresa, FechaEmision: ordenServicioModel.FechaEmision, EstadoOrdenServicio: estadoOrdenServicio, SubTotalConIVA: ordenServicioModel.SubTotalConIVA, SubTotalSinIVA: ordenServicioModel.SubTotalSinIVA, Tecnico: tecnico, Cliente: cliente, Descuento: ordenServicioModel.Descuento, ValorIVA: ordenServicioModel.ValorIVA, Total: ordenServicioModel.Total, Observaciones: ordenServicioModel.Observaciones, DiagnosticoRecepcion: ordenServicioModel.DiagnosticoRecepcion, DiagnosticoTecnico: ordenServicioModel.DiagnosticoTecnico}
}

func CreateOrdenServicio(c *fiber.Ctx) error {
	var ordenServicio models.OrdenServicio

	if err := c.BodyParser(&ordenServicio); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var tecnico models.Tecnico
	if err := FindTecnico(ordenServicio.TecnicoRefer, &tecnico); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var cliente models.Cliente
	if err := FindCliente(ordenServicio.ClienteRefer, &cliente); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var empresa models.Empresa
	if err := findEmpresa(ordenServicio.EmpresaRefer, &empresa); err != nil {
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

	var estadoOrdenServicio models.EstadoOrdenServicio
	if err := FindEstadoOrdenServicio(ordenServicio.EstadoOrdenServicioRefer, &estadoOrdenServicio); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&ordenServicio)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseTecnico := CreateResponseTecnico(tecnico, responseCiudad, responseEmpresa)
	responseCliente := CreateResponseCliente(cliente)
	responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estadoOrdenServicio, responseEmpresa)
	responseOrdenServicio := CreateResponseOrdenServicio(ordenServicio, responseEmpresa, responseEstadoOrdenServicio, responseTecnico, responseCliente)

	return c.Status(200).JSON(responseOrdenServicio)
}

func GetOrdenesServicio(c *fiber.Ctx) error {
	ordenesServicio := []models.OrdenServicio{}
	database.DB.Find(&ordenesServicio)
	responseOrdenesServicio := []OrdenServicio{}

	for _, ordenServicio := range ordenesServicio {

		var empresa models.Empresa
		database.DB.Find(&empresa, "id = ?", ordenServicio.EmpresaRefer)

		var ciudad models.Ciudad
		database.DB.Find(&ciudad, "id = ?", empresa.CiudadRefer)

		var provincia models.Provincia
		database.DB.Find(&provincia, "id = ?", ciudad.ProvinciaRefer)

		var tecnico models.Tecnico
		database.DB.Find(&tecnico, "id = ?", ordenServicio.TecnicoRefer)

		var cliente models.Cliente
		database.DB.Find(&cliente, "id = ?", ordenServicio.ClienteRefer)

		var estadoOrdenServicio models.EstadoOrdenServicio
		database.DB.Find(&estadoOrdenServicio, "id = ?", ordenServicio.EstadoOrdenServicioRefer)

		responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
		responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
		responseTecnico := CreateResponseTecnico(tecnico, responseCiudad, responseEmpresa)
		responseCliente := CreateResponseCliente(cliente)
		responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estadoOrdenServicio, responseEmpresa)
		responseOrdenServicio := CreateResponseOrdenServicio(ordenServicio, responseEmpresa, responseEstadoOrdenServicio, responseTecnico, responseCliente)
		responseOrdenesServicio = append(responseOrdenesServicio, responseOrdenServicio)
	}
	return c.Status(200).JSON(responseOrdenesServicio)
}

func FindOrdenServicio(id int, ordenServicio *models.OrdenServicio) error {
	database.DB.Find(&ordenServicio, "id = ?", id)
	if ordenServicio.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}

func GetOrdenServicio(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var ordenServicio models.OrdenServicio

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindOrdenServicio(id, &ordenServicio); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var empresa models.Empresa
	database.DB.First(&empresa, ordenServicio.EmpresaRefer)

	var ciudad models.Ciudad
	database.DB.First(&ciudad, empresa.CiudadRefer)

	var provincia models.Provincia
	database.DB.First(&provincia, ciudad.ProvinciaRefer)

	var tecnico models.Tecnico
	database.DB.First(&tecnico, ordenServicio.TecnicoRefer)

	var cliente models.Cliente
	database.DB.First(&cliente, ordenServicio.ClienteRefer)

	var estadoOrdenServicio models.EstadoOrdenServicio
	database.DB.First(&estadoOrdenServicio, ordenServicio.EstadoOrdenServicioRefer)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseTecnico := CreateResponseTecnico(tecnico, responseCiudad, responseEmpresa)
	responseCliente := CreateResponseCliente(cliente)
	responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estadoOrdenServicio, responseEmpresa)
	responseOrdenServicio := CreateResponseOrdenServicio(ordenServicio, responseEmpresa, responseEstadoOrdenServicio, responseTecnico, responseCliente)

	return c.Status(200).JSON(responseOrdenServicio)
}

func UpdatetOrdenServicio(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var ordenServicio models.OrdenServicio

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindOrdenServicio(id, &ordenServicio); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatetOrdenServicio struct {
		NumOrden                 string  `json:"numOrden"`
		EmpresaRefer             int     `json:"empresa_id"`
		FechaEmision             string  `json:"fecha_emision"`
		EstadoOrdenServicioRefer int     `json:"estado_orden_servicio_id"`
		SubTotalConIVA           float64 `json:"sub_total_con_IVA"`
		SubTotalSinIVA           float64 `json:"sub_total_sin_IVA"`
		TecnicoRefer             int     `json:"tecnico_id"`
		ClienteRefer             int     `json:"cliente_id"`
		Descuento                float64 `json:"descuento"`
		ValorIVA                 float64 `json:"valor_IVA"`
		Total                    float64 `json:"total"`
		Observaciones            string  `json:"observaciones"`
		DiagnosticoRecepcion     string  `json:"diagnostico_recepcion"`
		DiagnosticoTecnico       string  `json:"diagnostico_tecnico"`
	}

	var updateData UpdatetOrdenServicio

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	ordenServicio.NumOrden = updateData.NumOrden
	ordenServicio.EmpresaRefer = updateData.EmpresaRefer
	ordenServicio.FechaEmision = updateData.FechaEmision
	ordenServicio.EstadoOrdenServicioRefer = updateData.EstadoOrdenServicioRefer
	ordenServicio.SubTotalConIVA = updateData.SubTotalConIVA
	ordenServicio.SubTotalSinIVA = updateData.SubTotalSinIVA
	ordenServicio.TecnicoRefer = updateData.TecnicoRefer
	ordenServicio.ClienteRefer = updateData.ClienteRefer
	ordenServicio.Descuento = updateData.Descuento
	ordenServicio.ValorIVA = updateData.ValorIVA
	ordenServicio.Total = updateData.Total
	ordenServicio.Observaciones = updateData.Observaciones
	ordenServicio.DiagnosticoRecepcion = updateData.DiagnosticoRecepcion
	ordenServicio.DiagnosticoTecnico = updateData.DiagnosticoTecnico

	var empresa models.Empresa
	database.DB.First(&empresa, ordenServicio.EmpresaRefer)

	var ciudad models.Ciudad
	database.DB.First(&ciudad, empresa.CiudadRefer)

	var provincia models.Provincia
	database.DB.First(&provincia, ciudad.ProvinciaRefer)

	var tecnico models.Tecnico
	database.DB.First(&tecnico, ordenServicio.TecnicoRefer)

	var cliente models.Cliente
	database.DB.First(&tecnico, ordenServicio.ClienteRefer)

	var estadoOrdenServicio models.EstadoOrdenServicio
	database.DB.First(&estadoOrdenServicio, ordenServicio.EstadoOrdenServicioRefer)

	database.DB.Save(&ordenServicio)

	responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
	responseTecnico := CreateResponseTecnico(tecnico, responseCiudad, responseEmpresa)
	responseCliente := CreateResponseCliente(cliente)
	responseEstadoOrdenServicio := CreateResponseEstadoOrdenServicio(estadoOrdenServicio, responseEmpresa)
	responseOrdenServicio := CreateResponseOrdenServicio(ordenServicio, responseEmpresa, responseEstadoOrdenServicio, responseTecnico, responseCliente)

	return c.Status(200).JSON(responseOrdenServicio)
}
