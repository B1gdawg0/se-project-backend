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
		); 
		err != nil{ panic("Can't Automigrate Database: Auto shutdown...")}

	userRepo := repositories.ProvideUserRepository(db)
	userService := user.ProvideUserService(userRepo)
	userHandler := handlers.ProvideUserRestHandler(userService)

	authService := auth.ProvideAuthService(userRepo)
	authHandler := handlers.ProvideAuthRestHandler(authService)

	orderRepo := repositories.ProvideOrderReposity(db)
	orderService := order.ProvideOrderService(orderRepo)
	orderHandler := handlers.ProvideOrderRestHandler(orderService)

	oLineRepo := repositories.ProvideOrderLineRepository(db)
	oLineService := orderline.ProvideOrderLineService(oLineRepo)
	oLineHandler := handlers.ProvideOrderLineRestHandler(oLineService)

	tableRepo := repositories.ProvideTableRepository(db)
	tableService := table.ProvideTableService(tableRepo,orderRepo,oLineRepo)
	tableHandler := handlers.ProvideTableRestHandler(tableService)

	adminService := admin.ProvideAdminService(userRepo)
	adminHandler := handlers.ProvideAdminRestHandler(adminService)
	adminHandler.InitializeAdminAccount()

	handler := rest.ProvideHandler(userHandler, authHandler, tableHandler, orderHandler, oLineHandler, adminHandler)

	return handler
}