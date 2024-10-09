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

func (pg *PGGormDB) FindUserByID(id pgtype.UUID) (*entities.User, error)  {
	user := new(entities.User)

	if err := pg.db.Where("id = ?", id).Find(user); err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) FindUserByEmail(email string) (*entities.User, error)  {
	user := new(entities.User)

	if err := pg.db.Where("c_email = ?", email).Find(user); err.Error != nil{
		return nil, err.Error
	}

	return user, nil
}
func (pg *PGGormDB) FindUserByPhone(phone string) (*entities.User, error) {
	user := new(entities.User)

	if err := pg.db.Where("c_phone = ?", phone).Find(user); err.Error != nil{
		return nil, err.Error
	}

	return user, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) FindAll() ([]*entities.User, error) {
	users := new([]*entities.User)

	if err := pg.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return *users, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) CreateUser(rq *entities.User) (*response.CreateUserResponse, error) {

	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	response := &response.CreateUserResponse{
		ID: rq.ID,
		Name: rq.Name,
		Email: rq.Email,
		Password: rq.Password,
		Phone: rq.Phone,
	}	
		
	return response, nil
}


func (pg *PGGormDB) FindAllTable() ([]*entities.Table, error){
	var tables []*entities.Table

	if err := pg.db.Find(&tables).Error; err != nil {
		return nil, err
	}

	return tables, nil
}


func (pg *PGGormDB) CreateTable(rq *entities.Table) (*entities.Table, error){
	if err := pg.db.Create(rq).Error; err != nil{
		return nil, err
	}

	return rq, nil
}


func (pg *PGGormDB) FindTableByID(id string) (*entities.Table, error){
	table := new(entities.Table)
	if err := pg.db.Where("id = ?", id).Find(table).Error; err != nil{
		return nil, err
	}


	return table, nil
}

func (pg *PGGormDB) UpdateTableByID(rq *entities.Table) (*entities.Table, error){
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


func (pg *PGGormDB) DeleteTableByID(id string) (error){
	table:= new(entities.Table)

	if err := pg.db.First(table, "id = ?", id).Error; err != nil{
		return err
	}


	if err := pg.db.Delete(table).Error; err != nil{
		return err
	}


	return nil
}

func (pg *PGGormDB) FindAllOrder() ([]*entities.Order, error){
	orders := new([]*entities.Order)

	if err := pg.db.Find(orders).Error; err != nil{
		return nil, err
	}

	return *orders,nil
}

func (pg *PGGormDB) FindOrderByID(id pgtype.UUID) (*entities.Order, error){
	order := new(entities.Order)
	if err := pg.db.First(order, "id = ?", id).Error; err != nil{
		return nil, err
	}

	return order, nil
}


func (pg *PGGormDB) CreateOrderByID(rq *entities.Order) (*entities.Order, error){
	if err := pg.db.Create(rq).Error; err != nil{
		return nil, err
	}


	return rq, nil
}

func (pg *PGGormDB) DeleteOrderByID(id pgtype.UUID)(error){
	order := new(entities.Order)

	if err := pg.db.Where("id = ?",id).Find(order).Error; err != nil{
		return err
	}

	if err := pg.db.Delete(order).Error; err != nil{
		return err
	}


	return nil
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

func (pg *PGGormDB) DeleteOrderLineByID(id pgtype.UUID) (error){
	oline := new(entities.OrderLine)

	if err := pg.db.Where("id = ?",id).Find(oline).Error; err != nil{
		return err
	}

	if err := pg.db.Delete(oline).Error; err != nil{
		return err
	}


	return nil
}