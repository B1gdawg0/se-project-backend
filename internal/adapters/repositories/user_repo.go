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

func (uri *UserRepositoryImpl) FindUserByEmail(email string) (*entities.User, error) {
	return uri.Queries.FindUserByEmail(email)
}

func (uri *UserRepositoryImpl) FindUserByID(id pgtype.UUID) (*entities.User, error){
	return uri.Queries.FindUserByID(id)
}
func (uri *UserRepositoryImpl) FindUserByPhone(phone string) (*entities.User, error){
	return uri.Queries.FindUserByPhone(phone)
}

func (uri *UserRepositoryImpl) FindAll() ([]*entities.User, error) {
	return uri.Queries.FindAll()
}

func (uri *UserRepositoryImpl) Save(user *entities.User) (*response.CreateUserResponse, error) {
	response, err := uri.Queries.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return response, nil

}
