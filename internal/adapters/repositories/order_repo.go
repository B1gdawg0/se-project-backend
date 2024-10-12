package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/order"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type OrderReposityImpl struct {
	Queries queries.Database
}

func ProvideOrderReposity(db *gorm.DB) order.OrderReposity {
	return &OrderReposityImpl{
		Queries: queries.New(db),
	}
}

func (ori *OrderReposityImpl) FindAllOrder() ([]*entities.Order, error) {
	return ori.Queries.FindAllOrder()
}

func (ori *OrderReposityImpl) FindOrderByID(id pgtype.UUID) (*entities.Order, error) {
	return ori.Queries.FindOrderByID(id)
}

func (ori *OrderReposityImpl) CreateOrderByID(rq *entities.Order) (*entities.Order, error) {
	return ori.Queries.CreateOrderByID(rq)
}

func (ori *OrderReposityImpl) DeleteOrderByID(id pgtype.UUID) error {
	return ori.Queries.DeleteOrderByID(id)
}