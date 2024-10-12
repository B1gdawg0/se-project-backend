package user

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface{
	FindUserByID(id pgtype.UUID) (*response.FindUserResponse, error)
	FindUserByEmail(email string) (*response.FindUserResponse, error)
	FindAll() (*response.FindUsersResponse, error)
	Save(user *entities.User) (*response.CreateUserResponse, error)
}