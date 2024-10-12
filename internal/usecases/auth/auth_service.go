package auth

import (
	"errors"
	"os"
	"time"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface{
	Login(*requests.LoginRequest) (*response.LoginResponse, error)
	Register(rq *entities.User) (*response.RegisterResponse, error)
}


type AuthService struct{
	repo user.UserRepository
}

func ProvideAuthService(repo user.UserRepository) AuthUseCase{
	return &AuthService{
		repo: repo,
	}
}


func (as *AuthService) Login(rq *requests.LoginRequest) (*response.LoginResponse, error){
	user, err := as.repo.FindUserByEmail(rq.Email)

	if err != nil{
		return nil, err
	}

	if user.Email == "" || user.Password == ""{
		return nil, errors.New("email or password must not be empty string")
	}

	if user.Email != rq.Email {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rq.Password)); err != nil{
		return nil, err
	}

	token, err := generateJWT(user.Email, convert.UUIDToString(user.ID), "customer")

	if err != nil{
		return nil, err
	}

	response := &response.LoginResponse{
		Token: token,
	}

	return response , nil
}




func (uc *AuthService) Register(rq *entities.User) (*response.RegisterResponse, error){
	if rq.Name == "" && rq.Email == "" && rq.Password == "" && rq.Phone == ""{
		return nil, errors.New("value can't be empty string")
	}

	if user, _ := uc.repo.FindUserByEmail(rq.Email); user.Email == rq.Email{
		return nil, errors.New("this email is alreadys use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)

	if err != nil{
		return nil, err
	}

	rq.ID = utils.GenerateUUID()
	rq.Password = string(hashedPassword)

	user, err := uc.repo.Save(rq)

	if err != nil{
		return nil, err
	}

	response := &response.RegisterResponse{
		ID: convert.UUIDToString(user.ID),
		Name: user.Name,
		Email: user.Email,
	}

	return  response, nil
}


func generateJWT(email string, id string, role string) (string, error){
	skey := []byte(os.Getenv("JWT_SECRET_KEY"))
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":email,
		"user_id":id,
		"role": role,
		"exp":time.Now().Add(time.Hour * 48).Unix(),
	})

	tokenStr, err := token.SignedString(skey)

	if err != nil{
		return "", err
	}

	return tokenStr, nil
}