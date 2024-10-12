package user

import (
	"errors"

	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserUseCase interface {
	FindUserByID(id pgtype.UUID) (*response.FindUserResponse, error)
	FindUserByEmail(email string) (*response.FindUserResponse, error)
	FindAll() (*response.FindUsersResponse, error)
}

type UserService struct {
	repo UserRepository
}

func ProvideUserService(repo UserRepository) UserUseCase {
	return &UserService{
		repo: repo,
	}
}

func (uc *UserService) FindAll() (*response.FindUsersResponse, error) {
	list, err := uc.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (uc *UserService) FindUserByEmail(email string) (*response.FindUserResponse, error) {
	response, err := uc.repo.FindUserByEmail(email)

	if err != nil || response.Email == ""{
		return nil, errors.New("user not found")
	}

	return response, nil
}

func (uc *UserService) FindUserByID(id pgtype.UUID) (*response.FindUserResponse, error) {
	response, err := uc.repo.FindUserByID(id)

	if err != nil || response.Email == ""{
		return nil, errors.New("user not found")
	}

	return response, nil
}