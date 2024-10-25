package rest

import (
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest/handlers"
)

type Handler struct{
	User *handlers.UserRestHandler
	Auth *handlers.AuthRestHandler
	Table *handlers.TableRestHandler
	Order *handlers.OrderRestHandler
	Menu *handlers.MenuRestHandler
	OrderLine *handlers.OrderLineRestHandler
	Admin *handlers.AdminRestHandler
}


func ProvideHandler(user *handlers.UserRestHandler, auth *handlers.AuthRestHandler, table *handlers.TableRestHandler, order *handlers.OrderRestHandler, orderline *handlers.OrderLineRestHandler, admin *handlers.AdminRestHandler,  menu *handlers.MenuRestHandler) *Handler{
	return &Handler{
		User: user,
		Auth: auth,
		Table: table,
		Order: order,
		Menu: menu,
		OrderLine: orderline,
		Admin: admin,
	}
}