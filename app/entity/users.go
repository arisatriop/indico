package entity

type User struct {
	ID        int    `json:"-"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"-"`
	Roles     string `json:"roles"`
	CreatedAt string `json:"-" gorm:"-"`
	CreatedBy string `json:"-" gorm:"-"`
	UpdatedAt string `json:"-" gorm:"-"`
	UpdatedBy string `json:"-" gorm:"-"`
	DeletedAt string `json:"-" gorm:"-"`
	DeletedBy string `json:"-" gorm:"-"`
}
