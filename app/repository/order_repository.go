package repository

import (
	"indico-technical-test/app/entity"
	"indico-technical-test/database"
)

type IOrderRepository interface {
	FindAll() ([]entity.Order, error)
	FindByID(int) (*entity.Order, error)
}

type OrderRepository struct {
	DB *database.Postgres
}

func NewOrderRepository(db *database.Postgres) IOrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	var orders []entity.Order

	err := r.DB.GDB.Table("orders").Where("deleted_at IS NULL").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) FindByID(id int) (*entity.Order, error) {
	var order entity.Order

	err := r.DB.GDB.Table("orders").Where("id = ? AND deleted_at IS NULL", id).First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}
