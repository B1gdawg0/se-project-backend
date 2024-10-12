package response

type CreateTableResponse struct{
	ID string `json:"t_id"`
	C_ID string `json:"c_id"`
	Status string `json:"t_status"`
}

type GetTableResponse CreateTableResponse

type GetTablesResponse struct{
	Tables []GetTableResponse `json:"tables"`
}

type UpdateTableResponse struct{
	ID string `json:"t_id"`
	C_ID string `json:"c_id"`
	Status string `json:"t_status"`

	Order CreateOrderResponse `json:"order"`
	OrderLine CreateOrderLineResponse `json:"order-line"`
}