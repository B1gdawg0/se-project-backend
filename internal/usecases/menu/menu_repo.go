package menu

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
)

type MenuRepository interface {
	CreateMenu(rq *entities.Menu) (*entities.Menu, error)
	FindMenuByID(id string) (*entities.Menu, error)
	FindAllMenu() ([]*entities.Menu, error)
	UpdateMenu(rq *entities.Menu) (*entities.Menu, error)
	DeleteMenu(id string) error
}
