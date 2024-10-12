package response

type CreateOrderLineResponse struct{
	ID       string `json:"id"`
    Time     string `json:"time"`
    O_ID     string `json:"o_id"`
    M_ID     string `json:"m_id"`
    Quantity string `json:"l_quantity"`
    Price    string `json:"l_price"`
    Url      string `json:"l_urlslip"`
}


type GetOrderLineResponse CreateOrderLineResponse

type GetOrderLinesResponse struct{
	Olines []GetOrderLineResponse
}