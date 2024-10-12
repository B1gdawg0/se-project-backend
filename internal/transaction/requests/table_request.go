package requests


type CreateTableRequest struct{
	ID string `json:"t_id"`
}

type UpdateTableRequest struct{
	ID string `json:"id"`
	C_ID string `json:"c_id"`
	Status string `json:"t_status"`

	OrderLine CreateOrderLineRequest `json:"orderline"`
}