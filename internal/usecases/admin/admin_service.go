package admin

import (
	"errors"
	"github.com/B1gdawg0/se-project-backend/config"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase interface {
	InitializeAdminAccount () (*response.RegisterAdminResponse, error)
}

type AdminService struct {
	repo user.UserRepository
}

func ProvideAdminService(repo user.UserRepository) *AdminService {
    return &AdminService{
        repo: repo,
    }
}


func (uc *AdminService) InitializeAdminAccount() (*response.RegisterAdminResponse, error) {
	config := config.ProvideConfig()
	if (config.ADMIN_EMAIL == "" || config.ADMIN_PASSWORD == "") {
		return nil, errors.New("value can't be empty string")
	}

	if user, _ := uc.repo.FindUserByEmail(config.ADMIN_EMAIL); user.Email == config.ADMIN_EMAIL{
		return nil, errors.New("this email is alreadys use")
	}


	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.ADMIN_PASSWORD), bcrypt.DefaultCost)
	
	if err != nil{
		return nil, err
	}
	
	ID := utils.GenerateUUID()
	Password := string(hashedPassword)
	
	
	
	user_rq := entities.User{ID: ID, Email: config.ADMIN_EMAIL, Password: Password}

	user, err := uc.repo.Save(&user_rq)

	if err != nil{
		return nil, err
	}

	response := &response.RegisterAdminResponse{
		ID: convert.UUIDToString(user.ID),
		Email: user.Email,
	}

	return  response, nil
	


}