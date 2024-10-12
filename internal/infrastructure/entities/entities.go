package entities

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
    ID       pgtype.UUID `gorm:"primaryKey;type:uuid"`
    Email    string `gorm:"column:c_email"`
    Password string `gorm:"column:c_password"`
    Name     string `gorm:"column:c_name"`
    Phone    string `gorm:"column:c_phone"`
    
    MusicLines []MusicLine `gorm:"foreignKey:C_ID"`
    IGLines    []IGLine    `gorm:"foreignKey:C_ID"`
    Tables     []Table     `gorm:"foreignKey:C_ID"`
    Discount   []Discount  `gorm:"foreignKey:C_ID"`
}

type MusicLine struct {
    ID   pgtype.UUID `gorm:"primaryKey;type:uuid"`
    C_ID pgtype.UUID `gorm:"column:c_id"`
    Name string `gorm:"column:m_name"`
}

type IGLine struct {
    ID   pgtype.UUID `gorm:"primaryKey;type:uuid"`
    C_ID pgtype.UUID `gorm:"column:c_id"`
    Name string `gorm:"column:ig_account"`
}

type Table struct {
    ID     string `gorm:"primaryKey;type:string"`
    C_ID   pgtype.UUID `gorm:"column:c_id"`
    Status string   `gorm:"type:char(1);default:'A';check:t_status IN ('A', 'O', 'R');column:t_status"`
}

type Order struct {
    ID         pgtype.UUID `gorm:"primaryKey;type:uuid"`
    T_ID       string `gorm:"column:t_id"`
    Time       time.Time `gorm:"column:o_time"`

    Discount   Discount `gorm:"foreignKey:O_ID;references:ID"`
    Table      Table `gorm:"foreignKey:T_ID;references:ID"`
    OrderLines []OrderLine `gorm:"foreignKey:O_ID"`
}

type OrderLine struct {
    ID       pgtype.UUID `gorm:"primaryKey;type:uuid"`
    Time     time.Time   `gorm:"column:o_time"`
    O_ID     pgtype.UUID `gorm:"column:o_id"`
    M_ID     pgtype.UUID `gorm:"column:m_id"`
    Quantity int         `gorm:"column:l_quantity"`
    Price    float32     `gorm:"column:l_price"`
    Url      string      `gorm:"column:l_urlslip"`

    Menu Menu `gorm:"foreignKey:M_ID;references:ID"`
}

type Menu struct {
    ID          pgtype.UUID   `gorm:"primaryKey;type:uuid"`
    Price       float32  `gorm:"column:m_price"`
    Description string   `gorm:"column:m_description"`
    Url         []string `gorm:"type:text[]"`
}


type Discount struct {
    ID pgtype.UUID `gorm:"primaryKey;type:uuid"`
    C_ID pgtype.UUID `gorm:"column:c_id"`
    O_ID pgtype.UUID `gorm:"column:o_id"`
    Percent float32 `gorm:"column:d_percent"`
    Name string `gorm:"column:d_name"`
    Code string `gorm:"column:d_code"`
    Description string `gorm:"column:d_description"`
    StartDate time.Time `gorm:"column:d_start"`
    ExpDate time.Time `gorm:"column:d_exp"`
    Status bool `gorm:"column:d_status"`
}