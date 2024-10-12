package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Queries queries.Database
}

func ProvideUserRepository(db *gorm.DB) user.UserRepository {
	return &UserRepositoryImpl{
		Queries: queries.New(db),
	}
}

func (uri *UserRepositoryImpl) FindUserByEmail(email string) (*response.FindUserResponse, error){
	response, err := uri.Queries.FindUserByEmail(email)
	
	if err != nil{
		return nil, err
	}

	return response, nil
}

func (uri *UserRepositoryImpl) FindUserByID(id pgtype.UUID) (*response.FindUserResponse, error){
	response, err := uri.Queries.FindUserByID(id)
	
	if err != nil{
		return nil, err
	}

	return response, nil
}

func (uri *UserRepositoryImpl) FindAll() (*response.FindUsersResponse, error) {
	response, err := uri.Queries.FindAll()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (uri *UserRepositoryImpl) Save(user *entities.User) (*response.CreateUserResponse, error) {
	response, err := uri.Queries.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return response, nil

}
