package table

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/order"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/orderline"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type TableUseCase interface {
	FindAllTable() (*response.GetTablesResponse, error)
	FindTableByID(string) (*response.GetTableResponse, error)
	CreateTable(*entities.Table) (*response.CreateTableResponse, error)
	UpdateTableByID(*requests.UpdateTableRequest) (*response.UpdateTableResponse, error)
	DeleteTableByID(string) (error)
}

type TableService struct {
	tableRepo TableRepository
	orderRepo order.OrderReposity
	olineRepo orderline.OrderLineRepository
}

func ProvideTableService(table TableRepository, order order.OrderReposity, oline orderline.OrderLineRepository) TableUseCase {
	return &TableService{
		tableRepo: table,
		orderRepo: order,
		olineRepo: oline,
	}
}

func (ts *TableService) FindTableByID(id string) (*response.GetTableResponse, error){
	table, err := ts.tableRepo.FindTableByID(id)

	if err != nil {
        return nil, errors.New("failed to retrieve table: " + err.Error())
    }

    if table.ID == "" {
        return nil, errors.New("table not found")
    }

	C_IDStr := convert.UUIDToString(table.C_ID)
	if C_IDStr == ""{
		C_IDStr = "Nobody booked this table"
	}


	return &response.GetTableResponse{
		ID: table.ID,
		C_ID: C_IDStr,
		Status: table.Status,
	}, nil
}

func (ts *TableService) FindAllTable() (*response.GetTablesResponse, error){
	list, err := ts.tableRepo.FindAllTable()

	if err != nil{
		return nil, err
	}

	res := &response.GetTablesResponse{
		Tables: make([]response.GetTableResponse, 0),
	}

	for _, obj := range(list) {
		C_IDStr := convert.UUIDToString(obj.C_ID)

		if C_IDStr == ""{
		C_IDStr = "Nobody book this table"
		}

		res.Tables = append(res.Tables, response.GetTableResponse{
			ID: obj.ID,
			C_ID: C_IDStr,
			Status: obj.Status,
		})
	}

	return res, nil
}

func (ts *TableService) CreateTable(rq *entities.Table) (*response.CreateTableResponse, error) {
	selected, err := ts.tableRepo.FindTableByID(rq.ID)

	if err != nil{
		return nil, err
	}

	if selected != nil && selected.ID != ""{
		return nil, errors.New("table with this ID already exists")
	}

	table, err := ts.tableRepo.CreateTable(rq)


	if err != nil{
		return nil ,err
	}

	C_IDStr := convert.UUIDToString(table.C_ID)
	if C_IDStr == ""{
		C_IDStr = "Nobody book this table"
	}

	return &response.CreateTableResponse{
		ID: table.ID,
		C_ID: C_IDStr,
		Status: table.Status,
	}, nil
}

func (ts *TableService) UpdateTableByID(rq *requests.UpdateTableRequest) (*response.UpdateTableResponse, error) {
	exist, err := ts.tableRepo.FindTableByID(rq.ID)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if exist == nil {
		return nil, errors.New("table not found")
	}

	c_id, err := utils.StringToUUID(rq.C_ID)
	if err != nil {
		return nil, err
	}

	if exist.Status == "A" && rq.Status != "A" {
		orderPayload := &entities.Order{
			ID:   utils.GenerateUUID(),
			T_ID: exist.ID,
			Time: time.Now(),
		}

		order, err := ts.orderRepo.CreateOrderByID(orderPayload)
		if err != nil {
			return nil, err
		}

		m_id, quantity, price, err := ts.parseOrderLine(rq)
		if err != nil {
			ts.rollbackOrder(order.ID)
			return nil, err
		}

		olinePayload := &entities.OrderLine{
			ID:       utils.GenerateUUID(),
			Time:     order.Time,
			O_ID:     order.ID,
			M_ID:     *m_id,
			Quantity: quantity,
			Price:    float32(price),
			Url:      rq.OrderLine.Url,
		}

		oline, err := ts.olineRepo.CreateOrderLine(olinePayload)
		if err != nil {
			ts.rollbackOrder(order.ID)
			return nil, err
		}

		return ts.updateTableWithOrderNOrderLine(order, oline, exist, c_id, rq)
	}

	return ts.updateTable(rq, c_id)
}

func (ts *TableService) rollbackOrder(id pgtype.UUID) {
	if err := ts.orderRepo.DeleteOrderByID(id); err != nil {
		panic("Can't roll back order in booking table")
	}
}

func (ts *TableService) rollbackOrderLine(id pgtype.UUID) {
	if err := ts.olineRepo.DeleteOrderLineByID(id); err != nil {
		panic("Can't roll back order in booking table")
	}
}

func (ts *TableService) parseOrderLine(rq *requests.UpdateTableRequest) (*pgtype.UUID, int, float64, error) {
	m_id, err := utils.StringToUUID(rq.OrderLine.M_ID)
	if err != nil {
		return nil, 0, 0, err
	}

	quantity, err := strconv.Atoi(rq.OrderLine.Quantity)
	if err != nil {
		return nil, 0, 0, err
	}

	price, err := strconv.ParseFloat(rq.OrderLine.Price, 64)
	if err != nil {
		return nil, 0, 0, err
	}

	return m_id, quantity, price, nil
}

func (ts *TableService) updateTableWithOrderNOrderLine(order *entities.Order, oline *entities.OrderLine, exist *entities.Table, c_id *pgtype.UUID, rq *requests.UpdateTableRequest) (*response.UpdateTableResponse, error) {
	C_IDStr := convert.UUIDToString(exist.C_ID)
	if C_IDStr == "" {
		C_IDStr = "Nobody booked this table"
	}

	tablePayload := &entities.Table{
		ID:     rq.ID,
		C_ID:   *c_id,
		Status: rq.Status,
	}

	table, err := ts.tableRepo.UpdateTableByID(tablePayload)
	if err != nil {
		ts.rollbackOrder(order.ID)
		ts.rollbackOrderLine(oline.ID)
		return nil, err
	}

	return &response.UpdateTableResponse{
		ID:     table.ID,
		C_ID:   C_IDStr,
		Status: table.Status,
		Order: response.CreateOrderResponse{
			ID:   convert.UUIDToString(order.ID),
			T_ID: order.T_ID,
			Time: order.Time.Format("2006-01-02 15:04:05"),
		},
		OrderLine: response.CreateOrderLineResponse{
			ID:       convert.UUIDToString(oline.ID),
			Time:     order.Time.Format("2006-01-02 15:04:05"),
			O_ID:     convert.UUIDToString(oline.O_ID),
			M_ID:     convert.UUIDToString(oline.M_ID),
			Quantity: fmt.Sprintf("%d", oline.Quantity),
			Price:    fmt.Sprintf("%.2f", oline.Price),
			Url:      oline.Url,
		},
	}, nil
}

func (ts *TableService) updateTable(rq *requests.UpdateTableRequest, c_id *pgtype.UUID) (*response.UpdateTableResponse, error) {
	tablePayload := &entities.Table{
		ID:     rq.ID,
		C_ID:   *c_id,
		Status: rq.Status,
	}

	table, err := ts.tableRepo.UpdateTableByID(tablePayload)
	if err != nil {
		return nil, err
	}

	C_IDStr := convert.UUIDToString(table.C_ID)
	if C_IDStr == "" {
		C_IDStr = "Nobody booked this table"
	}

	return &response.UpdateTableResponse{
		ID:     table.ID,
		C_ID:   C_IDStr,
		Status: table.Status,
	}, nil
}




func (ts *TableService) DeleteTableByID(id string) error{
	 exist, err := ts.tableRepo.FindTableByID(id)

	 if err != nil{
		return err
	 }

	 if exist == nil {
		return errors.New("table not found")
	 }

	 return ts.tableRepo.DeleteTableByID(id)
}
