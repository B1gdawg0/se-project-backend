package requests

type CreateMenuRequest struct {
	Price       float64 `json:"m_price"`
	Description string  `json:"m_description"`
	Url         string  `json:"url"`
}

type UpdateMenuRequest struct {
	Price       float64 `json:"m_price"`
	Description string  `json:"m_description"`
	Url         string  `json:"url"`
}
