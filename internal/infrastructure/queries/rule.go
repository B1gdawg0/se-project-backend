package queries

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/jackc/pgx/v5/pgtype"
)

type Database interface{
	CreateUser(rq *entities.User) (*response.CreateUserResponse, error)
	FindUserByEmail(email string) (*entities.User, error) 
	FindUserByPhone(phone string) (*entities.User, error) 
	FindUserByID(id pgtype.UUID) (*entities.User, error) 
	FindAll() ([]*entities.User, error) 


	CreateTable(rq *entities.Table) (*entities.Table, error)
	FindAllTable()([]*entities.Table, error)
	FindTableByID(id string) (*entities.Table, error)
	UpdateTableByID(rq *entities.Table) (*entities.Table, error)
	DeleteTableByID(id string) ( error)

	FindOrderByID(id pgtype.UUID) (*entities.Order, error)
	FindAllOrder()([]*entities.Order,error)
	CreateOrderByID(rq *entities.Order) (*entities.Order, error)
	DeleteOrderByID(id pgtype.UUID) (error)

	FindOrderLineByID(id pgtype.UUID) (*entities.OrderLine, error)
	FindAllOrderLine()([]*entities.OrderLine, error)
	CreateOrderLine(rq *entities.OrderLine) (*entities.OrderLine, error)
} 