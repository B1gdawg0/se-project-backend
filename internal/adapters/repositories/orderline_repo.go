package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/orderline"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type OrderLineRepositoryImpl struct {
	Queries queries.Database
}

func ProvideOrderLineRepository(db *gorm.DB) orderline.OrderLineRepository {
	return &OrderLineRepositoryImpl{
		Queries: queries.New(db),
	}
}

func (olr *OrderLineRepositoryImpl) FindAllOrderLine() ([]*entities.OrderLine, error) {
	return olr.Queries.FindAllOrderLine()
}

func (olr *OrderLineRepositoryImpl) FindOrderLineByID(id pgtype.UUID) (*entities.OrderLine, error) {
	return olr.Queries.FindOrderLineByID(id)
}

func (olr *OrderLineRepositoryImpl) CreateOrderLine(rq *entities.OrderLine) (*entities.OrderLine, error) {
	return olr.Queries.CreateOrderLine(rq)
}

func (olr *OrderLineRepositoryImpl) DeleteOrderLineByID(id pgtype.UUID) error {
	return olr.Queries.DeleteOrderByID(id)
}
