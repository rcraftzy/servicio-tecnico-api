package controllers 

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/models"
)

type Producto struct {
  ID           int `json:"id"`
  Codigo       string `json:"codigo"`
  Nombre       string `json:"nombre"`
  PrecioVenta       float64 `json:"precioVenta"`
  StockMin       float64 `json:"stockMin"`
  StockMax       float64 `json:"stockMax"`
  Stock       float64 `json:"stock"`
  ControlaStock bool `json:"controlaStock"`
  AplicaIVA bool `json:"aplicaIva"`
  Empresa Empresa `json:"empresa"`
}

func CreateResponseProducto(producto models.Producto, empresa Empresa) Producto {
  return Producto {ID: producto.ID ,Codigo: producto.Codigo ,Nombre: producto.Nombre, PrecioVenta: producto.PrecioVenta, StockMin: producto.StockMin, StockMax: producto.StockMax, Stock: producto.Stock, ControlaStock: producto.ControlaStock, AplicaIVA: producto.AplicaIVA, Empresa: empresa}
}

func CreateProducto(c *fiber.Ctx) error {
	var producto models.Producto

	if err := c.BodyParser(&producto); err != nil {
		return c.Status(400).JSON(err.Error())
	}

  var empresa models.Empresa
	if err := findEmpresa(producto.EmpresaRefer, &empresa); err != nil {
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

	database.DB.Create(&producto)

  responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
  responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
  responseProducto := CreateResponseProducto(producto, responseEmpresa)

	return c.Status(200).JSON(responseProducto)
}

func GetProductos(c *fiber.Ctx) error {
	productos := []models.Producto{}
	database.DB.Find(&productos)
	responseProductos := []Producto{}

	for _, producto := range productos {

    var empresa models.Empresa
		database.DB.Find(&empresa, "id = ?", producto.EmpresaRefer)

    var ciudad models.Ciudad
		database.DB.Find(&ciudad, "id = ?", empresa.CiudadRefer)

		var provincia models.Provincia
		database.DB.Find(&provincia, "id = ?", ciudad.ProvinciaRefer)

    responseCiudad := CreateResponseCiudad(ciudad, CreateResponseProvincia(provincia))
		responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
    responseProducto := CreateResponseProducto(producto, responseEmpresa)
    responseProductos = append(responseProductos, responseProducto)
	}
	return c.Status(200).JSON(responseProductos)
}

func FindProducto(id int, producto *models.Producto) error {
	database.DB.Find(&producto, "id = ?", id)
	if producto.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}

func GetProducto(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var producto models.Producto

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindProducto(id, &producto); err != nil {
		return c.Status(400).JSON(err.Error())
	}

  var empresa models.Empresa
	database.DB.First(&empresa, producto.EmpresaRefer)
  
	var ciudad models.Ciudad
	database.DB.First(&ciudad, empresa.CiudadRefer)

	var provincia models.Provincia
	database.DB.First(&provincia, ciudad.ProvinciaRefer)

  responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
  responseProducto := CreateResponseProducto(producto, responseEmpresa)

	return c.Status(200).JSON(responseProducto)
}

func UpdateProducto(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var producto models.Producto

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindProducto(id, &producto); err != nil {
		return c.Status(400).JSON(err.Error())
	}

  type UpdateProducto struct {
    Codigo       string `json:"codigo"`
    Nombre       string `json:"nombre"`
    PrecioVenta       float64 `json:"precioVenta"`
    StockMin       float64 `json:"stockMin"`
    StockMax       float64 `json:"stockMax"`
    Stock       float64 `json:"stock"`
    ControlaStock bool `json:"controlaStock"`
    AplicaIVA bool `json:"aplicaIva"`
    EmpresaRefer int `json:"empresa_id"`
  }

	var updateData UpdateProducto

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	producto.Codigo = updateData.Codigo
	producto.Nombre = updateData.Nombre
	producto.PrecioVenta = updateData.PrecioVenta
	producto.StockMin = updateData.StockMin
	producto.StockMax = updateData.StockMax
	producto.Stock = updateData.Stock
	producto.ControlaStock = updateData.ControlaStock
	producto.AplicaIVA = updateData.AplicaIVA
	producto.EmpresaRefer = updateData.EmpresaRefer

  var empresa models.Empresa
	if err := findEmpresa(producto.EmpresaRefer, &empresa); err != nil {
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
	database.DB.Save(&producto)

  responseProvincia := CreateResponseProvincia(provincia)
	responseCiudad := CreateResponseCiudad(ciudad, responseProvincia)
	responseEmpresa := CreateResponseEmpresa(empresa, responseCiudad)
  responseProducto := CreateResponseProducto(producto, responseEmpresa)
	return c.Status(200).JSON(responseProducto)
}
