package entity

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SKU        string `json:"sku"`
	Quantity   int    `json:"quantity"`
	LocationID string `json:"localtionId"`
	CreatedAt  string `json:"-"`
	CreatedBy  string `json:"-"`
	UpdatedAt  string `json:"-"`
	UpdatedBy  string `json:"-"`
	DeletedAt  string `json:"-"`
	DeletedBy  string `json:"-"`
}
