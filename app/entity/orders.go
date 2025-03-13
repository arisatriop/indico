package entity

type Order struct {
	ID          int    `json:"id"`
	OrderNumber string `json:"orderNumber"`
	OrderType   string `json:"orderType"`
	ProductID   int    `json:"productId"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	CreatedAt   string `json:"-"`
	CreatedBy   string `json:"-"`
	UpdatedAt   string `json:"-"`
	UpdatedBy   string `json:"-"`
	DeletedAt   string `json:"-"`
	DeletedBy   string `json:"-"`
}
