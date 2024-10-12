package main

import (
	"github.com/B1gdawg0/se-project-backend/internal/wire"
	"github.com/B1gdawg0/se-project-backend/router"
	"github.com/gofiber/fiber/v2"
)

func main(){
	handler := wire.InitHandler()

	app := fiber.New()

	router.RegisterApiRouter(app, handler)


	app.Listen(":8000")
}