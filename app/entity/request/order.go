package request

type OrderRequest struct {
	ProductID int `json:"productId" form:"productId" validate:"required"`
	Quantity  int `json:"quantity" form:"quantity" validate:"required"`
}
