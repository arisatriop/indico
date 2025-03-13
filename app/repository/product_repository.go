package repository

import (
	"indico-technical-test/app/entity"
	"indico-technical-test/database"
	"time"
)

type IProductRepository interface {
	Create(*entity.Product) error
	Update(*entity.Product) error
	Delete(id int) error
	FindAll() ([]entity.Product, error)
	FindByID(int) (*entity.Product, error)
}

type ProductRepository struct {
	DB *database.Postgres
}

func NewProductRepository(db *database.Postgres) IProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Create(p *entity.Product) error {

	db := r.DB.GDB.Table("products")

	if err := db.Omit("UpdatedAt", "CreatedAt", "DeletedAt").Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Update(p *entity.Product) error {
	db := r.DB.GDB.Table("products").Where("id = ?", p.ID)

	if p.CreatedAt == "" {
		db = db.Omit("CreatedAt")
	}
	if p.UpdatedAt == "" {
		db = db.Omit("UpdatedAt")
	}
	if p.DeletedAt == "" {
		db = db.Omit("DeletedAt")
	}

	if err := db.Updates(p).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Delete(id int) error {

	err := r.DB.GDB.Table("products").Where("id = ?", id).Update("deleted_at", time.Now().Format("2006-01-02 15:04:05.000")).Error
	if err != nil {
		return err
	}

	return nil

}

func (r *ProductRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.DB.GDB.Table("products").Where("deleted_at is null").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) FindByID(id int) (*entity.Product, error) {
	var product entity.Product
	if err := r.DB.GDB.Table("products").Where("id = ? and deleted_at is null", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
