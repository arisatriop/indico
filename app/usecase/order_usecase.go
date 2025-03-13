package usecase

import (
	"errors"
	"fmt"
	"indico-technical-test/app/entity"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/app/repository"
	"indico-technical-test/database"
	"strings"
	"sync"
	"time"

	"math/rand"

	"gorm.io/gorm"
)

type IOrderUsecase interface {
	FindAll() ([]entity.Order, error)
	FindByID(int) (*entity.Order, error)
	Receive([]request.OrderRequest) error
	Ship([]request.OrderRequest) error
}

type OrderUsecase struct {
	DB                *database.Postgres
	OrderRepository   repository.IOrderRepository
	ProductRepository repository.IProductRepository
}

func NewOrderUsecase(db *database.Postgres, orderRepo repository.IOrderRepository, productRepo repository.IProductRepository) IOrderUsecase {
	return &OrderUsecase{
		DB:                db,
		OrderRepository:   orderRepo,
		ProductRepository: productRepo,
	}
}

func (u *OrderUsecase) FindAll() ([]entity.Order, error) {
	return u.OrderRepository.FindAll()
}

func (u *OrderUsecase) FindByID(id int) (*entity.Order, error) {

	order, err := u.OrderRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return order, nil
}

func (u *OrderUsecase) Receive(req []request.OrderRequest) error {

	wg := sync.WaitGroup{}
	errChan := make(chan error, len(req))

	for _, r := range req {
		wg.Add(1)

		go func(r request.OrderRequest) {
			defer wg.Done()

			trx := u.DB.GDB.Begin()
			product, err := u.ProductRepository.FindByID(r.ProductID)
			if err != nil {
				trx.Rollback()
				errChan <- fmt.Errorf("error find product: %v", err)
				return
			}

			err = trx.Table("products").Where("id = ?", r.ProductID).Update("quantity", product.Quantity+r.Quantity).Error
			if err != nil {
				trx.Rollback()
				errChan <- fmt.Errorf("error update product: %v", err)
				return
			}

			order := entity.Order{
				ProductID:   product.ID,
				ProductName: product.Name,
				Quantity:    r.Quantity,
				OrderType:   "Receive",
				OrderNumber: u.GenerateOrderNumber(),
			}

			err = trx.Table("orders").Omit("CreatedAt", "UpdatedAt", "DeletedAt").Create(&order).Error
			if err != nil {
				trx.Rollback()
				errChan <- fmt.Errorf("error create order: %v", err)
				return
			}

			trx.Commit()
		}(r)
	}

	wg.Wait()
	close(errChan)
	return nil
}

func (u *OrderUsecase) Ship(req []request.OrderRequest) error {

	wg := sync.WaitGroup{}
	errChan := make(chan error, len(req))

	for _, r := range req {
		wg.Add(1)

		go func(r request.OrderRequest) {
			defer wg.Done()

			trx := u.DB.GDB.Begin()
			product, err := u.ProductRepository.FindByID(r.ProductID)
			if err != nil {
				trx.Rollback()
				errChan <- fmt.Errorf("error find product: %v", err)
				return
			}

			if product.Quantity < r.Quantity {
				trx.Rollback()
				errChan <- fmt.Errorf("product quantity is not enough")
				return
			}

			err = trx.Table("products").Where("id = ?", r.ProductID).Update("quantity", product.Quantity-r.Quantity).Error
			if err != nil {
				trx.Rollback()
				errChan <- fmt.Errorf("error update product: %v", err)
				return
			}

			order := entity.Order{
				ProductID:   product.ID,
				ProductName: product.Name,
				Quantity:    r.Quantity,
				OrderType:   "Ship",
				OrderNumber: u.GenerateOrderNumber(),
			}

			err = trx.Table("orders").Omit("CreatedAt", "UpdatedAt", "DeletedAt").Create(&order).Error
			if err != nil {
				trx.Rollback()
				errChan <- fmt.Errorf("error create order: %v", err)
				return
			}

			trx.Commit()
		}(r)
	}

	wg.Wait()
	close(errChan)
	return nil
}

func (u *OrderUsecase) GenerateOrderNumber() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	orderNumber := strings.ReplaceAll(now.Format("20060102150405.000000"), ".", "")
	return fmt.Sprintf("%s%d", orderNumber, rand.Intn(90)+10)
}
