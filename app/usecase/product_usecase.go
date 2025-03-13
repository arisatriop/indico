package usecase

import (
	"errors"
	"indico-technical-test/app/entity"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/app/repository"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type IProductUsecase interface {
	Create(*request.ProductCreateRequest) error
	Update(*request.ProductUpdateRequest) error
	Delete(int) error
	FindAll() ([]entity.Product, error)
	FindByID(int) (*entity.Product, error)
}

type ProductUsecase struct {
	ProductRepository repository.IProductRepository
}

func NewProductUsecase(productRepo repository.IProductRepository) IProductUsecase {
	return &ProductUsecase{
		ProductRepository: productRepo,
	}
}

func (u *ProductUsecase) Create(req *request.ProductCreateRequest) error {

	qty, _ := strconv.Atoi(req.Quantity)

	product := entity.Product{
		Name:       req.Name,
		SKU:        req.SKU,
		Quantity:   qty,
		LocationID: req.LocationID,
	}

	return u.ProductRepository.Create(&product)
}

func (u *ProductUsecase) Update(req *request.ProductUpdateRequest) error {

	qty, _ := strconv.Atoi(req.Quantity)

	product := entity.Product{
		ID:         req.ID,
		Name:       req.Name,
		SKU:        req.SKU,
		Quantity:   qty,
		LocationID: req.LocationID,
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05.000"),
	}

	return u.ProductRepository.Update(&product)
}

func (u *ProductUsecase) Delete(id int) error {
	return u.ProductRepository.Delete(id)
}

func (u *ProductUsecase) FindAll() ([]entity.Product, error) {
	return u.ProductRepository.FindAll()
}

func (u *ProductUsecase) FindByID(id int) (*entity.Product, error) {

	product, err := u.ProductRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return product, nil
}
