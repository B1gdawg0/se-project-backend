package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/table"
	"gorm.io/gorm"
)


type TableRepositoryImpl struct{
	Queries queries.Database
}



func ProvideTableRepository(db *gorm.DB) table.TableRepository{
	return &TableRepositoryImpl{
		Queries: queries.New(db),
	}
}

func (tri *TableRepositoryImpl) FindTableByID(id string) (*entities.Table, error){
	return tri.Queries.FindTableByID(id)
} 

func (tri *TableRepositoryImpl) FindAllTable() ([]*entities.Table, error){
	return tri.Queries.FindAllTable()
}

func (tri *TableRepositoryImpl) CreateTable(rq *entities.Table) (*entities.Table, error){
	return tri.Queries.CreateTable(rq)
}

func(tri *TableRepositoryImpl) UpdateTableByID(rq *entities.Table) (*entities.Table, error){
	return tri.Queries.UpdateTableByID(rq)
}

func (tri *TableRepositoryImpl) DeleteTableByID(id string) (error){
	return tri.Queries.DeleteTableByID(id)
}