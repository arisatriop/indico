package request

type LocationRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Capacity int    `json:"capacity" form:"capacity" validate:"required"`
}
