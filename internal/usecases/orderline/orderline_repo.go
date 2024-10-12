package orderline

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderLineRepository interface{
	FindOrderLineByID(id pgtype.UUID) (*entities.OrderLine, error)
	FindAllOrderLine()([]*entities.OrderLine, error)
	CreateOrderLine(rq *entities.OrderLine) (*entities.OrderLine, error)
	DeleteOrderLineByID(id pgtype.UUID) (error)
}