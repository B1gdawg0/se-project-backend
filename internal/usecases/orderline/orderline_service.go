package orderline

import (
	"errors"
	"fmt"
	"time"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type OrderLineUseCase interface {
	FindOrderLineByID(id pgtype.UUID) (*response.GetOrderLineResponse, error)
	FindAllOrderLine() (*response.GetOrderLinesResponse, error)
	CreateOrderLine(*entities.OrderLine) (*response.CreateOrderLineResponse, error)
}

type OrderLineService struct {
	repo OrderLineRepository
}

func ProvideOrderLineService(repo OrderLineRepository) OrderLineUseCase {
	return &OrderLineService{
		repo: repo,
	}
}

func (ols *OrderLineService) FindAllOrderLine() (*response.GetOrderLinesResponse, error) {
	all, err := ols.repo.FindAllOrderLine()

	if err != nil{
		return nil, err
	}

	list := response.GetOrderLinesResponse{
		Olines: make([]response.GetOrderLineResponse, 0),
	}


	for _, obj := range(all) {
		list.Olines = append(list.Olines, response.GetOrderLineResponse{
			ID:       convert.UUIDToString(obj.ID),
			Time:     obj.Time.Format("2006-01-02 15:04:05"),
			O_ID:     convert.UUIDToString(obj.O_ID),
			M_ID:     convert.UUIDToString(obj.M_ID),
			Quantity: fmt.Sprintf("%d", obj.Quantity),
			Price:    fmt.Sprintf("%.2f", obj.Price),
			Url:      obj.Url,
		})
	}


	return &list, nil
}

func (ols *OrderLineService) FindOrderLineByID(id pgtype.UUID) (*response.GetOrderLineResponse, error) {
	orderline,err := ols.repo.FindOrderLineByID(id)

	if err != nil{
		return nil ,err
	}

	return &response.GetOrderLineResponse{
		ID:       convert.UUIDToString(orderline.ID),
		Time:     orderline.Time.Format("2006-01-02 15:04:05"),
		O_ID:     convert.UUIDToString(orderline.O_ID),
		M_ID:     convert.UUIDToString(orderline.M_ID),
		Quantity: fmt.Sprintf("%d", orderline.Quantity),
		Price:    fmt.Sprintf("%.2f", orderline.Price),
		Url:      orderline.Url,
	},nil
}

func (ols *OrderLineService) CreateOrderLine(rq *entities.OrderLine) (*response.CreateOrderLineResponse, error) {
	rq.ID = utils.GenerateUUID()

	selected, err := ols.repo.FindOrderLineByID(rq.ID)

	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, err
	}

	if selected != nil && selected.ID == rq.ID{
		return nil, errors.New("order already exists")
	}

	rq.Time = time.Now()

	orderline, err := ols.repo.CreateOrderLine(rq)

	if err != nil {
		return nil, err
	}

	return &response.CreateOrderLineResponse{
		ID:       convert.UUIDToString(orderline.ID),
		Time:     orderline.Time.Format("2006-01-02 15:04:05"),
		O_ID:     convert.UUIDToString(orderline.O_ID),
		M_ID:     convert.UUIDToString(orderline.M_ID),
		Quantity: fmt.Sprintf("%d", orderline.Quantity),
		Price:    fmt.Sprintf("%.2f", orderline.Price),
		Url:      orderline.Url,
	}, nil
}
