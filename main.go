package main

import (
	"github.com/B1gdawg0/se-project-backend/background"
	"github.com/B1gdawg0/se-project-backend/internal/wire"
	"github.com/B1gdawg0/se-project-backend/rabbit"
	"github.com/B1gdawg0/se-project-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	handler := wire.InitHandler()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	router.RegisterApiRouter(app, handler)

	rabbit.InitRabbitMQConnection()
	// rabbit.TestEmailQueueConnection()
	defer rabbit.CleanupRabbitMQ()

	rabbit.InitEmailWorker()

	go background.DailyMonitor(app, handler)

	app.Listen(":8000")
}