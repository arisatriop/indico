package request

type ProductCreateRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	SKU        string `json:"sku" form:"sku" validate:"required"`
	Quantity   string `json:"quantity" form:"quantity" validate:"required"`
	LocationID string `json:"locationId" form:"locationId" validate:"required"`
}

type ProductUpdateRequest struct {
	ID         int    `json:"id" form:"id" validate:"required"`
	Name       string `json:"name" form:"name" validate:"required"`
	SKU        string `json:"sku" form:"sku" validate:"required"`
	Quantity   string `json:"quantity" form:"quantity" validate:"required"`
	LocationID string `json:"locationId" form:"locationId" validate:"required"`
}
