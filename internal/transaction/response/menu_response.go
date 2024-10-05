package response

import "github.com/jackc/pgx/v5/pgtype"

type CreateMenuResponse struct {
	ID          pgtype.UUID `gorm:"primaryKey;type:uuid"`
	Price       float64     `gorm:"column:m_price"`
	Description string      `gorm:"column:m_description"`
	Url         string    `gorm:"type:text"`
}

type GetMenuResponse CreateMenuResponse

type GetAllMenuResponse struct {
		Menu []GetMenuResponse `json:"Menu"`
}

type UpdateMenuResponse CreateMenuResponse