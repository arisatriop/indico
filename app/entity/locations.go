package entity

type Location struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	CreatedAt string `json:"-" gorm:"-"`
	CreatedBy string `json:"-" gorm:"default:null"`
	UpdatedAt string `json:"-" gorm:"-"`
	UpdatedBy string `json:"-" gorm:"default:null"`
	DeletedAt string `json:"-" gorm:"default:null"`
	DeletedBy string `json:"-" gorm:"default:null"`
}
