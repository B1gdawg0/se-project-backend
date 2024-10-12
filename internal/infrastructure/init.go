package infrastructure

import (
	"fmt"

	"github.com/B1gdawg0/se-project-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvidePostGresDB(config config.Config) *gorm.DB{
	postgresURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDB)

	db, err := gorm.Open(postgres.Open(postgresURI), &gorm.Config{})

	if err != nil{
		panic("Set up Postgres err: "+err.Error())
	}

	fmt.Println("PostGres Database Start...")

	return db
}