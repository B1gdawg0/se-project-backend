package user

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface{
	FindUserByID(id pgtype.UUID) (*entities.User, error) 
	FindUserByEmail(email string) (*entities.User, error) 
	FindUserByPhone(phone string) (*entities.User, error)
	FindAll() ([]*entities.User, error) 
	Save(user *entities.User) (*response.CreateUserResponse, error)
}