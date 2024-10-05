package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/menu"
	"gorm.io/gorm"
)

type MenuRepositoryImpl struct {
	Queries queries.Database
}

func ProvideMenuRepository(db *gorm.DB) menu.MenuRepository {
	return &MenuRepositoryImpl{
		Queries: queries.New(db),
	}
}

func (mri *MenuRepositoryImpl) CreateMenu(rq *entities.Menu) (*entities.Menu, error) {
	return mri.Queries.CreateMenu(rq)
}

func (mri *MenuRepositoryImpl) FindMenuByID(id string) (*entities.Menu, error) {
	return mri.Queries.FindMenuByID(id)
}

func (mri *MenuRepositoryImpl) FindAllMenu() ([]*entities.Menu, error) {
	return mri.Queries.FindAllMenu()
}

func (mri *MenuRepositoryImpl) UpdateMenu(rq *entities.Menu) (*entities.Menu, error) {
	return mri.Queries.UpdateMenu(rq)
}

func (mri *MenuRepositoryImpl) DeleteMenu(id string) error {
	return mri.Queries.DeleteMenu(id)
}
