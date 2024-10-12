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
	orderLine := app.Group("/order-lines")
	admin := app.Group("/admin")

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

	
	orderLine.Get("",rqHandler.OrderLine.GetOrderLines)
	orderLine.Get("/id=:id",rqHandler.OrderLine.GetOrderLineByID)
	orderLine.Post("", rqHandler.OrderLine.CreateOrderLine)

	admin.Get("/users", rqHandler.User.GetUsers)
	admin.Get("/users/id=:id", rqHandler.User.GetUserByID)
	admin.Get("/users/email=:email", rqHandler.User.GetUserByEmail)

	admin.Get("/tables", rqHandler.Table.GetTables)
	admin.Get("/tables/id=:id", rqHandler.Table.GetTableByID)

	admin.Get("/orders", rqHandler.Order.GetAllOrder)
	admin.Get("/orders/id=:id", rqHandler.Order.GetOrderByID)

	admin.Get("/order-lines", rqHandler.OrderLine.GetOrderLines)
	admin.Get("/order-lines/id=:id", rqHandler.OrderLine.GetOrderLineByID)

}