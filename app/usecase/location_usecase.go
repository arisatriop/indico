package usecase

import (
	"indico-technical-test/app/entity"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/app/repository"
)

type ILocationUsecase interface {
	Create(*request.LocationRequest) error
	FindAll() ([]entity.Location, error)
}

type LocationUsecase struct {
	LocationRepository repository.ILocationRepository
}

func NewLocationUsecase(locationRepo repository.ILocationRepository) ILocationUsecase {
	return &LocationUsecase{
		LocationRepository: locationRepo,
	}
}

func (u *LocationUsecase) Create(req *request.LocationRequest) error {
	return u.LocationRepository.Create(req)
}

func (u *LocationUsecase) FindAll() ([]entity.Location, error) {
	return u.LocationRepository.FindAll()
}
