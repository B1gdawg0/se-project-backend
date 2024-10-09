package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gofiber/fiber/v2"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/utils"

	
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
func (urh *UserRestHandler) GetCustomerByPhone(c *fiber.Ctx) error{
	phone := c.Params("phone")

	if phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
		})
	}

	response, err := urh.usecase.FindUserByPhone(phone)

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
func (trh *UserRestHandler) CreateUser(c *fiber.Ctx) error{
	rq := new(requests.CreateUserRequest)

	if err := c.BodyParser(rq); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":err.Error(),
		})
	}

	payload := &entities.User{
		ID: utils.GenerateUUID(),
		Name: rq.Name,
		Email: rq.Email,
		Password: rq.Password,
		Phone: rq.Phone,
	}


	response, err := trh.usecase.Save(payload)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Bad Request",
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful create User",
		"payload":fiber.Map{
			"table":response,
		},
	})
}