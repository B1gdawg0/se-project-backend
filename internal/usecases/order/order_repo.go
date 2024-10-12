package order

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/jackc/pgx/v5/pgtype"
)


type OrderReposity interface{
	FindAllOrder()([]*entities.Order, error)
	FindOrderByID(id pgtype.UUID) (*entities.Order, error)
	CreateOrderByID(*entities.Order) (*entities.Order, error)
	DeleteOrderByID(id pgtype.UUID)(error)
}