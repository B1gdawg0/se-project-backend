package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthRestHandler struct{
	usecase auth.AuthUseCase
}

func ProvideAuthRestHandler(usecase auth.AuthUseCase) *AuthRestHandler{
	return &AuthRestHandler{
		usecase: usecase,
	}
}


func (arh *AuthRestHandler) Login(c *fiber.Ctx) error{
	req := new(requests.LoginRequest)

	if err := c.BodyParser(req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":err,
		})
	}

	response, err := arh.usecase.Login(req)

	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Fail login!",
			"payload":fiber.Map{
				"status":false,
			},
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful login!",
		"payload":fiber.Map{
			"status":response,
		},
	})

}


func (arh *AuthRestHandler) Register(c *fiber.Ctx) error{
	rq := new(requests.RegisterRequest)

	if err:= c.BodyParser(rq); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":err.Error(),
		})
	}

	payload := &entities.User{
		Name: rq.Name,
		Email: rq.Email,
		Password: rq.Password,
		Phone: rq.Phone,
	}

	user, err := arh.usecase.Register(payload)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error":err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Successful Created",
		"payload":fiber.Map{
			"user":user,
		},
	})
}
