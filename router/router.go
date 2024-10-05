package router

import (
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest"
	"github.com/B1gdawg0/se-project-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRouter(app *fiber.App, rqHandler *rest.Handler){
	user := app.Group("/users")
	auth := app.Group("/auth")
	table := app.Group("/tables")
	order := app.Group("/orders")
	menu := app.Group("/menu")
	orderLine := app.Group("/order-lines")

	auth.Post("/login", rqHandler.Auth.Login)
	auth.Post("/register", rqHandler.Auth.Register)

	user.Use(middleware.CheckJWT)
	user.Get("", rqHandler.User.GetUsers)
	user.Get("/id=:id", rqHandler.User.GetUserByID)
	user.Get("/email=:email",rqHandler.User.GetUserByEmail)


	table.Get("", rqHandler.Table.GetTables)
	table.Get("/id=:id", rqHandler.Table.GetTableByID)
	table.Post("", rqHandler.Table.CreateTable)
	table.Put("id=:id",rqHandler.Table.UpdateTableByID)
	table.Delete("/id=:id", rqHandler.Table.DeleteTableByID)

	order.Get("",rqHandler.Order.GetAllOrder)
	order.Get("/id=:id",rqHandler.Order.GetOrderByID)
	order.Post("",rqHandler.Order.CreateOrderByID)

	menu.Get("", rqHandler.Menu.GetAllMenu)
	menu.Get("/id=:id", rqHandler.Menu.GetMenuByID)
	menu.Post("", rqHandler.Menu.CreateMenu)
	menu.Put("/id=:id", rqHandler.Menu.UpdateMenuByID)
	menu.Delete("/id=:id", rqHandler.Menu.DeleteMenuByID)
  
	orderLine.Get("",rqHandler.OrderLine.GetOrderLines)
	orderLine.Get("/id=:id",rqHandler.OrderLine.GetOrderLineByID)
	orderLine.Post("", rqHandler.OrderLine.CreateOrderLine)
}