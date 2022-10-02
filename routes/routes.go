package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roberto-carlos-tg/go-auht/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
  // Provincia
  app.Get("/api/v1.0/provincia", controllers.GetProvincias)
  app.Post("/api/v1.0/provincia", controllers.CreateProvincia)
  app.Get("/api/v1.0/provincia/:id", controllers.GetProvincia)
  app.Delete("/api/v1.0/provincia/:id", controllers.DeleteProvincia)
  app.Put("/api/v1.0/provincia/:id", controllers.UpdateProvincia)
  // Ciudad
  app.Get("/api/v1.0/ciudad", controllers.GetCiudades)
  app.Get("/api/v1.0/ciudad/:id", controllers.GetCiudad)
  app.Post("/api/v1.0/ciudad", controllers.CreateCiudad)
  app.Put("/api/v1.0/ciudad/:id", controllers.UpdateCiudad)
  app.Delete("/api/v1.0/ciudad/:id", controllers.DeleteCiudad)
  // Empresa 
  app.Get("/api/v1.0/empresa", controllers.GetEmpresas)
  app.Get("/api/v1.0/empresa/:id", controllers.GetEmpresa)
  app.Post("/api/v1.0/empresa", controllers.CreateEmpresa)
  app.Put("/api/v1.0/empresa/:id", controllers.UpdateEmpresa)
  // Tecnico
  app.Get("/api/v1.0/tecnico", controllers.GetTecnicos)
  app.Get("/api/v1.0/tecnico/:id", controllers.GetTecnico)
  app.Post("/api/v1.0/tecnico", controllers.CreateTecnico)
  // Producto
  app.Get("/api/v1.0/producto", controllers.GetProductos)
  app.Get("/api/v1.0/producto/:id", controllers.GetProducto)
  app.Post("/api/v1.0/producto", controllers.CreateProducto)
  app.Put("/api/v1.0/producto/:id", controllers.UpdateProducto)
  // Cliente 
  app.Get("/api/v1.0/clientes", controllers.GetClientes)
  app.Post("/api/v1.0/clientes", controllers.CreateCliente)
  app.Get("/api/v1.0/clientes/:id", controllers.GetCliente)
  app.Put("/api/v1.0/clientes/:id", controllers.UpdateCliente)
  // Estado Orden Servicio
  app.Get("/api/v1.0/estadoordenservicio", controllers.GetEstadosOrdenServicio)
  app.Post("/api/v1.0/estadoordenservicio", controllers.CreateEstadoOrdenServicio)
  // app.Get("/api/v1.0/estadoordenservicio/:id", controllers.GetEstadoOrdenServicio)
  app.Get("/api/v1.0/estadoordenservicio/:id", controllers.GetEstadoOrdenServicioByEmpresa)
  app.Put("/api/v1.0/estadoordenservicio/:id", controllers.UpdateEstadoOrdenServicio)
  app.Delete("/api/v1.0/estadoordenservicio/:id", controllers.DeleteEstadoOrdenServicio)
  // Orden Servicio
  app.Get("/api/v1.0/ordenServicio", controllers.GetOrdenesServicio)
  app.Post("/api/v1.0/ordenServicio", controllers.CreateOrdenServicio)
  app.Get("/api/v1.0/ordenServicio/:id", controllers.GetOrdenServicio)
  app.Put("/api/v1.0/ordenServicio/:id", controllers.UpdatetOrdenServicio)
  // Detalle Orden Servicio
  app.Get("/api/v1.0/detalleordenServicio", controllers.GetDetalleOrdenesServicio)
  app.Post("/api/v1.0/detalleordenServicio", controllers.CreateDetalleOrdenesServicio)
  // app.Get("/api/v1.0/detalleordenServicio/:id", controllers.GetDetalleOrdenServicio)
  app.Get("/api/v1.0/detalleordenServicio/:orden_servicio_id", controllers.GetDetalleOrdenServicioByOrder)
  // Empresa User
  app.Get("/api/v1.0/empresaUser", controllers.GetUsersEmpresas)
  app.Get("/api/v1.0/empresaUser/:id", controllers.GetUserEmpresa)
  app.Post("/api/v1.0/empresaUser", controllers.CreateUserEmpresa)
  // Role
  app.Get("/api/v1.0/role", controllers.GetRoles)
  app.Get("/api/v1.0/role/:id", controllers.GetRole)
  app.Post("/api/v1.0/role", controllers.CreateRole)
  app.Put("/api/v1.0/role/:id", controllers.UpdateRole)
  // Role User
  app.Get("/api/v1.0/roleUser", controllers.GetRolesUser)
  app.Get("/api/v1.0/roleUser/:id", controllers.GetRoleUser)
  app.Post("/api/v1.0/roleUser", controllers.CreateRoleUser)
  app.Put("/api/v1.0/roleUser/:id", controllers.UpdateRoleUser)
}
