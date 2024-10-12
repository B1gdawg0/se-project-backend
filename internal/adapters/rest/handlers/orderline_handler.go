package handlers

import (
	"fmt"
	"strconv"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/orderline"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type OrderLineRestHandler struct{
	usecase orderline.OrderLineUseCase
}

func ProvideOrderLineRestHandler(usecase orderline.OrderLineUseCase) *OrderLineRestHandler{
	return &OrderLineRestHandler{
		usecase: usecase,
	}
}

func (olrh *OrderLineRestHandler) GetOrderLines(c *fiber.Ctx) error{
	list, err := olrh.usecase.FindAllOrderLine()

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful get order lines",
		"payload":fiber.Map{
			"order-lines":list,
		},
	})
}

func (olrh *OrderLineRestHandler) GetOrderLineByID(c *fiber.Ctx) error{
	id, err := utils.StringToUUID(c.Params("id"))

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad request",
			"error":err.Error(),
		})
	}

	response, err := olrh.usecase.FindOrderLineByID(*id)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful create order line",
		"payload":fiber.Map{
			"orderline":response,
		},
	})
}

func (olrh *OrderLineRestHandler) CreateOrderLine(c *fiber.Ctx) error{
	rq := new(requests.CreateOrderLineRequest)

	if err := c.BodyParser(rq); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad request",
			"error":err.Error(),
		})
	}

	o_id, err := utils.StringToUUID(rq.O_ID)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad request",
			"error":err.Error(),
		})
	}
	m_id, err := utils.StringToUUID(rq.M_ID)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad request",
			"error":err.Error(),
		})
	}
	quantity, err := strconv.Atoi(rq.Quantity)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad request",
			"error":err.Error(),
		})
	}
	price, err := strconv.ParseFloat(rq.Price, 32)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad request",
			"error":err.Error(),
		})
	}

	fmt.Println(*m_id)

	payload := entities.OrderLine{
		O_ID: *o_id,
		M_ID: *m_id,
		Quantity: quantity,
		Price: float32(price),
	}

	response, err := olrh.usecase.CreateOrderLine(&payload)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful create order line",
		"payload":fiber.Map{
			"orderline":response,
		},
	})
}