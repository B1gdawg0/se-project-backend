package order

import (
	"errors"
	"time"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)


type OrderUseCase interface{
	FindAllOrder()(*response.GetOrdersResponse, error)
	FindOrderByID(id pgtype.UUID) (*response.GetOrderResponse, error)
	CreateOrderByID(*entities.Order) (*response.CreateOrderResponse, error)
}

type OrderService struct{
	repo OrderReposity
}


func ProvideOrderService(repo OrderReposity) OrderUseCase{
	return &OrderService{
		repo: repo,
	}
}


func (os *OrderService) FindAllOrder()(*response.GetOrdersResponse, error){
	orders, err := os.repo.FindAllOrder()

	if err != nil {
		return nil, err
	}

	list := &response.GetOrdersResponse{
		Orders: make([]response.GetOrderResponse, 0),
	}

	for _, obj := range(orders) {
		list.Orders = append(list.Orders, response.GetOrderResponse{
			ID: convert.UUIDToString(obj.ID),
			T_ID: obj.T_ID,
			Time: obj.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return list,nil
}


func (os *OrderService) FindOrderByID(id pgtype.UUID) (*response.GetOrderResponse, error){
	selected, err := os.repo.FindOrderByID(id)

	if err != nil{
		return nil, err
	}

	return &response.GetOrderResponse{
		ID: convert.UUIDToString(selected.ID),
		T_ID: selected.T_ID,
		Time: selected.Time.Format("2006-01-02 15:04:05"),
	},nil
}

func (os *OrderService) CreateOrderByID(rq *entities.Order) (*response.CreateOrderResponse, error){
	rq.ID = utils.GenerateUUID()

	selected, err := os.repo.FindOrderByID(rq.ID)

	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, err
	}

	if selected != nil && selected.ID == rq.ID{
		return nil, errors.New("order already exists")
	}
	
	rq.Time = time.Now()
	order, err := os.repo.CreateOrderByID(rq)

	if err != nil{
		return nil, err
	}

	return &response.CreateOrderResponse{
		ID: convert.UUIDToString(order.ID),
		T_ID: order.T_ID,
		Time: order.Time.Format("2006-01-02 15:04:05"),
	}, nil
}
