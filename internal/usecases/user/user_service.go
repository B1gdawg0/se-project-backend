package user

import (
	"errors"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
)

type UserUseCase interface {
	FindUserByID(id pgtype.UUID) (*entities.User, error) 
	FindUserByEmail(email string) (*entities.User, error) 
	FindUserByPhone(phone string) (*entities.User, error)
	FindAll() ([]*entities.User, error) 
	Save(user *entities.User) (*response.CreateUserResponse, error)

}

type UserService struct {
	repo UserRepository
}

func ProvideUserService(repo UserRepository) UserUseCase {
	return &UserService{
		repo: repo,
	}
}

func (uc *UserService) FindAll() ([]*entities.User, error)  {
	list, err := uc.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (uc *UserService) FindUserByEmail(email string) (*entities.User, error)  {
	response, err := uc.repo.FindUserByEmail(email)

	if err != nil || response.Email == ""{
		return nil, errors.New("user not found")
	}

	return response, nil
}

func (uc *UserService) FindUserByID(id pgtype.UUID) (*entities.User, error) {
	response, err := uc.repo.FindUserByID(id)

	if err != nil || response.Email == ""{
		return nil, errors.New("user not found")
	}

	return response, nil
}
func (uc *UserService) FindUserByPhone(phone string) (*entities.User, error) {
	response, err := uc.repo.FindUserByPhone(phone)

	if err != nil || response.Phone == ""{
		return nil, errors.New("user not found")
	}

	return response, nil
}

func (uc *UserService) Save(user *entities.User) (*response.CreateUserResponse, error) {
    // ตรวจสอบว่า email มี "@___.com" หรือไม่
    if !utils.IsValidEmail(user.Email) {
        return nil, errors.New("invalid email format, must contain '@___.com'")
    }

    // ตรวจสอบว่า password มีครบ 8 ตัวหรือไม่
    if len(user.Password) < 8 {
        return nil, errors.New("password must be at least 8 characters long")
    }

    // ตรวจสอบว่า phone มี 10 ตัวหรือไม่
    if len(user.Phone) != 10 {
        return nil, errors.New("phone number must be exactly 10 digits")
    }

    selected, err := uc.repo.FindUserByID(user.ID)
    if err != nil {
        return nil, err
    }

    if selected != nil && selected.ID == user.ID {
        return nil, errors.New("user ID already exists")
    }

    selected, err1 := uc.repo.FindUserByEmail(user.Email)
    if err1 != nil {
        return nil, err1
    }

    if selected != nil && selected.Email == user.Email {
        return nil, errors.New("user Email already exists")
    }

    selected, err2 := uc.repo.FindUserByPhone(user.Phone)
    if err2 != nil {
        return nil, err2
    }
    
    if selected != nil && selected.Phone == user.Phone {
        return nil, errors.New("user Phone already exists")
    }

    user1, err := uc.repo.Save(user)
    if err != nil {
        return nil, err
    }

    return &response.CreateUserResponse{
        ID:       user1.ID,
        Email:    user1.Email,
        Name:     user1.Name,
        Phone:    user1.Phone,
        Password: user1.Password,
    }, nil
}

