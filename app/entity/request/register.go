package request

type RegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Roles    string `json:"roles" form:"roles" validate:"required"`
}
