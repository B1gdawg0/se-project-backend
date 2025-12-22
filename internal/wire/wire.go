package wire

import (
	"github.com/B1gdawg0/se-project-backend/config"
	"github.com/B1gdawg0/se-project-backend/internal/adapters/repositories"
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest"
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest/handlers"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/admin"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/auth"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/discount"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/igline"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/menu"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/musicline"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/order"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/orderline"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/table"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
)

func InitHandler() *rest.Handler{
	config := config.ProvideConfig()
	db := infrastructure.ProvidePostGresDB(*config)

	if err := db.AutoMigrate(
		&entities.User{},
        &entities.MusicLine{},
        &entities.IGLine{},
        &entities.Table{},
        &entities.Order{},
        &entities.OrderLine{},
        &entities.Menu{},
		&entities.Discount{},
		); 
		err != nil{ panic("Can't Automigrate Database: Auto shutdown...")
	}

	userRepo := repositories.ProvideUserRepository(db)
	userService := user.ProvideUserService(userRepo)
	userHandler := handlers.ProvideUserRestHandler(userService)

	authService := auth.ProvideAuthService(userRepo)
	authHandler := handlers.ProvideAuthRestHandler(authService)

	menuRepo := repositories.ProvideMenuRepository(db)
	menuService := menu.ProvideMenuService(menuRepo)
	menuHandler := handlers.ProvideMenuRestHandler(menuService)
  
	oLineRepo := repositories.ProvideOrderLineRepository(db)
	oLineService := orderline.ProvideOrderLineService(oLineRepo)
	oLineHandler := handlers.ProvideOrderLineRestHandler(oLineService)

	orderRepo := repositories.ProvideOrderReposity(db)
	orderService := order.ProvideOrderService(orderRepo,oLineRepo)
	orderHandler := handlers.ProvideOrderRestHandler(orderService)

	tableRepo := repositories.ProvideTableRepository(db)
	tableService := table.ProvideTableService(tableRepo,orderRepo,oLineRepo)
	tableHandler := handlers.ProvideTableRestHandler(tableService)

	adminService := admin.ProvideAdminService(userRepo)
	adminHandler := handlers.ProvideAdminRestHandler(adminService)
	adminHandler.InitializeAdminAccount()

	igLineRepo := repositories.ProvideIgLineRepository(db)
	igLineService := igline.ProvideIgLineService(igLineRepo)
	igLineHandler := handlers.ProvideIglineHandler(igLineService)

	musicLineRepo := repositories.ProvideMusicLineRepository(db)
	musicLineService := musicline.ProvideMusicService(musicLineRepo)
	musicLineHandler := handlers.ProvideMusicLineHandler(musicLineService)


	discountRepo := repositories.ProvideDiscountRepository(db)
	discountService := discount.ProvideDiscountService(discountRepo)
	discountHandler := handlers.ProvideDiscountRestHandler(discountService)


	handler := rest.ProvideHandler(userHandler, authHandler, tableHandler, orderHandler, oLineHandler, adminHandler, menuHandler, igLineHandler, musicLineHandler, discountHandler)

	return handler
}