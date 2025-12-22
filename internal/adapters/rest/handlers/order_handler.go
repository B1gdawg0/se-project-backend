package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/order"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/B1gdawg0/se-project-backend/rabbit"
	"github.com/gofiber/fiber/v2"
)


type OrderRestHandler struct{
	usecase order.OrderUseCase
}


func ProvideOrderRestHandler(usecase order.OrderUseCase) *OrderRestHandler{
	return &OrderRestHandler{
		usecase: usecase,
	}
}

func (orh *OrderRestHandler) GetAllOrder(c *fiber.Ctx) error{
	orders, err := orh.usecase.FindAllOrder()

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"succesful get all order",
		"payload":fiber.Map{
			"orders": orders,
		},
	})
}

func (orh *OrderRestHandler) GetOrderByID(c *fiber.Ctx) error{
	idStr := c.Params("id")

	uuid, err := utils.StringToUUID(idStr)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":"Invalid ID format",
		})
	}
	
	response, err := orh.usecase.FindOrderByID(*uuid)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful create order",
		"payload":fiber.Map{
			"order":response,
		},
	})
}

func (orh *OrderRestHandler) GetOrderByTableID(c *fiber.Ctx) error{
	idStr := c.Params("id")
	
	response, err := orh.usecase.FindOrderByTableID(idStr)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful create order",
		"payload":fiber.Map{
			"order":response,
		},
	})
}


func (orh *OrderRestHandler) CreateOrderByID(c *fiber.Ctx) error{
	rq := new(requests.CreateOrderRequest)

	if err := c.BodyParser(rq); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error": err.Error(),
		})
	}
	
	payload := &entities.Order{
		T_ID: rq.T_ID,
	}

	response, err := orh.usecase.CreateOrderByID(payload)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	id, err := utils.StringToUUID(response.ID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	data, err := orh.usecase.FindOrderByID(*id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	rabbit.SendMessage("order","email", data)

	return c.JSON(fiber.Map{
		"message":"Succesful create order",
		"payload":fiber.Map{
			"order":response,
		},
	})
}


func (orh *OrderRestHandler) CreateOrderWithOrderLines(c *fiber.Ctx) error{
	rq := new(requests.CreateOrderWithOrderLinesRequest)

	if err := c.BodyParser(rq);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	res, err := orh.usecase.CreateOrderWithOrderLines(rq)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := utils.StringToUUID(res.ID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	data, err := orh.usecase.FindOrderByID(*id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	rabbit.SendMessage("order","email", data)

	return c.JSON(res)
}