package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gofiber/fiber/v2"
)

type UserRestHandler struct{
	usecase user.UserUseCase
}

func ProvideUserRestHandler(usecase user.UserUseCase) *UserRestHandler{
	return &UserRestHandler{
		usecase: usecase,
	}
}

func (urh *UserRestHandler) GetUsers(c *fiber.Ctx) error{
	list, err := urh.usecase.FindAll()

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message":"Successful get list of user",
		"payload":fiber.Map{
			"users": list,
		},
	})
}


func (urh *UserRestHandler) GetUserByID(c *fiber.Ctx) error{
	id := convert.StringToUUID(c.Params("id"))

	response, err := urh.usecase.FindUserByID(id)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message":"Successful get list of user",
		"payload":fiber.Map{
			"user": response,
		},
	})
}

func (urh *UserRestHandler) GetUserByEmail(c *fiber.Ctx) error{
	email := c.Params("email")

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
		})
	}

	response, err := urh.usecase.FindUserByEmail(email)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message":"Successful get list of user",
		"payload":fiber.Map{
			"user": response,
		},
	})
}