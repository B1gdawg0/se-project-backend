package menu

import (
	"errors"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type MenuUseCase interface {
	CreateMenu(*entities.Menu) (*response.CreateMenuResponse, error)
	FindMenuByID(string) (*response.GetMenuResponse, error)
	FindAllMenu() (*response.GetAllMenuResponse, error)
	UpdateMenu(*entities.Menu) (*response.UpdateMenuResponse, error)
	DeleteMenu(string) error
}

type MenuService struct {
	repo MenuRepository
}

// ProvideMenuService creates a new instance of MenuService
func ProvideMenuService(repo MenuRepository) MenuUseCase {
	return &MenuService{
		repo: repo,
	}
}

// CreateMenu creates a new menu item
func (ms *MenuService) CreateMenu(rq *entities.Menu) (*response.CreateMenuResponse, error) {

	menu, err := ms.repo.CreateMenu(rq)
	if err != nil {
		return nil, err
	}

	return &response.CreateMenuResponse{
		ID:          menu.ID,
		Price:       menu.Price,
		Description: menu.Description,
		Url:         menu.Url,
	}, nil
}

func (ms *MenuService) FindMenuByID(id string) (*response.GetMenuResponse, error) {

	if id == "" {
		return nil, errors.New("menu ID is required")
	}

	menu, err := ms.repo.FindMenuByID(id)
	if err != nil {
		return nil, err
	}
	if menu == nil {
		return nil, errors.New("menu not found")
	}

	// Return the found menu response
	return &response.GetMenuResponse{
		ID:          menu.ID,
		Price:       menu.Price,
		Description: menu.Description,
		Url:         menu.Url,
	}, nil
}

// FindAllMenu retrieves all available menu items
func (ms *MenuService) FindAllMenu() (*response.GetAllMenuResponse, error) {
	menus, err := ms.repo.FindAllMenu()
	if err != nil {
		return nil, err
	}

	// Map entities to response objects
	var menuResponses []response.GetMenuResponse
	for _, menu := range menus {
		menuResponses = append(menuResponses, response.GetMenuResponse{
			ID:          menu.ID,
			Price:       menu.Price,
			Description: menu.Description,
			Url:         menu.Url,
		})
	}

	// Return the list of menus
	return &response.GetAllMenuResponse{
		Menu: menuResponses,
	}, nil
}

// UpdateMenu updates an existing menu item
func (ms *MenuService) UpdateMenu(rq *entities.Menu) (*response.UpdateMenuResponse, error) {
	zeroUUID := pgtype.UUID{}

	if rq.ID.Bytes == zeroUUID.Bytes {
		return nil, errors.New("menu ID is required")
	}

	menuID := uuid.UUID(rq.ID.Bytes).String() // Convert to uuid.UUID and then to string

	// Optionally, find the menu by ID to verify it exists before updating
	_, err := ms.FindMenuByID(menuID)
	if err != nil {
		return nil, errors.New("menu not found")
	}

	// Call the repository to update the menu
	updatedMenu, err := ms.repo.UpdateMenu(rq)
	if err != nil {
		return nil, err
	}

	// Return the updated menu response
	return &response.UpdateMenuResponse{
		ID:          updatedMenu.ID,
		Price:       updatedMenu.Price,
		Description: updatedMenu.Description,
		Url:         updatedMenu.Url,
	}, nil
}

func (ms *MenuService) DeleteMenu(id string) error {
	if id == "" {
		return errors.New("menu ID is required")
	}
	err := ms.repo.DeleteMenu(id)
	if err != nil {
		return err
	}

	return nil
}
