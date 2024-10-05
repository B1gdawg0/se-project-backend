package queries

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type PGGormDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) Database {
	return &PGGormDB{
		db: db,
	}
}

func (pg *PGGormDB) FindUserByID(id pgtype.UUID) (*response.FindUserResponse, error) {
	user := new(entities.User)

	if err := pg.db.Where("id = ?", id).Find(user); err.Error != nil {
		return nil, err.Error
	}

	response := &response.FindUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	return response, nil
}

func (pg *PGGormDB) FindAllTable() ([]*entities.Table, error) {
	var tables []*entities.Table

	if err := pg.db.Find(&tables).Error; err != nil {
		return nil, err
	}

	return tables, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) FindUserByEmail(email string) (*response.FindUserResponse, error) {
	user := new(entities.User)

	if err := pg.db.Where("c_email = ?", email).Find(user); err.Error != nil {
		return nil, err.Error
	}

	response := &response.FindUserResponse{
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
		Phone: user.Phone,
	}

	return response, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) FindAll() (*response.FindUsersResponse, error) {
	users := new([]*entities.User)

	if err := pg.db.Find(users).Error; err != nil {
		return nil, err
	}

	list := make([]response.FindUserResponse, len(*users))
	for i, ctx := range *users {
		list[i] = response.FindUserResponse{
			ID: ctx.ID,
			Name: ctx.Name,
			Email: ctx.Email,
			Phone: ctx.Phone,
		}
	}

	usersResponse := &response.FindUsersResponse{
		Users: list,
	}

	return usersResponse, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) CreateUser(rq *entities.User) (*response.CreateUserResponse, error) {

	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	response := &response.CreateUserResponse{
		ID:       rq.ID,
		Name:     rq.Name,
		Email:    rq.Email,
		Password: rq.Password,
		Phone:    rq.Phone,
	}

	return response, nil
}


func (pg *PGGormDB) FindAllTable() ([]*entities.Table, error){
	var tables []*entities.Table

	if err := pg.db.Find(&tables).Error; err != nil {
		return nil, err
	}

	return tables,nil
}


func (pg *PGGormDB) CreateTable(rq *entities.Table) (*entities.Table, error){
	if err := pg.db.Create(rq).Error; err != nil{
		return nil, err
	}

	return rq, nil
}

func (pg *PGGormDB) FindTableByID(id string) (*entities.Table, error) {
	table := new(entities.Table)
	if err := pg.db.Where("id = ?", id).Find(table).Error; err != nil {
		return nil, err
	}

	return table, nil
}

func (pg *PGGormDB) UpdateTableByID(rq *entities.Table) (*entities.Table, error) {
	table := new(entities.Table)
	if err := pg.db.First(table, "id = ?", rq.ID).Error; err != nil {
		return nil, err
	}

	table.C_ID = rq.C_ID
	table.Status = rq.Status

	if err := pg.db.Save(table).Error; err != nil {
		return nil, err
	}

	return table, nil
}

func (pg *PGGormDB) DeleteTableByID(id string) error {
	table := new(entities.Table)

	if err := pg.db.First(table, "id = ?", id).Error; err != nil {
		return err
	}

	if err := pg.db.Delete(table).Error; err != nil {
		return err
	}

	return nil
}

func (pg *PGGormDB) FindAllOrder() ([]*entities.Order, error) {
	orders := new([]*entities.Order)

	if err := pg.db.Find(orders).Error; err != nil {
		return nil, err
	}

	return *orders, nil
}

func (pg *PGGormDB) FindOrderByID(id pgtype.UUID) (*entities.Order, error) {
	order := new(entities.Order)
	if err := pg.db.First(order, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (pg *PGGormDB) CreateOrderByID(rq *entities.Order) (*entities.Order, error) {
	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	return rq, nil
}

func (pg *PGGormDB) FindAllOrderLine()([]*entities.OrderLine, error){
	olines := new([]*entities.OrderLine)

	if err := pg.db.Find(olines).Error; err!= nil{
		return nil, err
	}

	

	return *olines, nil
}


func (pg *PGGormDB) FindOrderLineByID(id pgtype.UUID) (*entities.OrderLine, error){
	oline := new(entities.OrderLine)
	if err := pg.db.First(oline, "id = ?",id).Error; err != nil{
		return nil, err
	}
	return oline, nil
}


func (pg *PGGormDB) CreateOrderLine(rq *entities.OrderLine) (*entities.OrderLine, error){
	if err := pg.db.Create(rq).Error; err != nil{
		return nil, err
	}

	return rq, nil
}

func (pg *PGGormDB) CreateMenu(rq *entities.Menu) (*entities.Menu, error) {
	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}
	return rq, nil
}

func (pg *PGGormDB) FindMenuByID(id string) (*entities.Menu, error) {
	Menu := new(entities.Menu)
	if err := pg.db.Where("id = ?", id).Find(Menu).Error; err != nil {
		return nil, err
	}
	return Menu, nil
}

func (pg *PGGormDB) FindAllMenu() ([]*entities.Menu, error) {
	var menuList []*entities.Menu

	if err := pg.db.Find(&menuList).Error; err != nil {
		return nil, err
	}

	return menuList, nil
}

func (pg *PGGormDB) UpdateMenu(rq *entities.Menu) (*entities.Menu, error) {
	menu := new(entities.Menu)

	if err := pg.db.First(menu, "id = ?", menu.ID).Error; err != nil {
		return nil, err
	}
	menu.Description = rq.Description
	menu.Price = rq.Price
	menu.Url = rq.Url

	if err := pg.db.Save(menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}

func (pg *PGGormDB) DeleteMenu(id string) error {
	menu := new(entities.Menu)
	if err := pg.db.First(menu, "id = ?", id).Error; err != nil {
		return err
	}

	if err := pg.db.Delete(menu).Error; err != nil {
		return err
	}

	return nil
}
