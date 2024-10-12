package table

import "github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"


type TableRepository interface{
	FindAllTable()([]*entities.Table, error)
	FindTableByID(id string) (*entities.Table, error)
	CreateTable(rq *entities.Table) (*entities.Table, error)
	UpdateTableByID(rq *entities.Table) (*entities.Table, error)
	DeleteTableByID(id string)(error)
}