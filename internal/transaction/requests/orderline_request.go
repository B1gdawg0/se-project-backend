package requests

type CreateOrderLineRequest struct{
    O_ID     string `json:"o_id"`
    M_ID     string `json:"m_id"`
    Quantity string    `json:"quantity"`
    Price    string `json:"price"`
    Url      string    `json:"urlslip"`
}