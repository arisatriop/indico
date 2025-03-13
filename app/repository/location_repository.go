package repository

import (
	"indico-technical-test/app/entity"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/database"
)

type ILocationRepository interface {
	Create(*request.LocationRequest) error
	FindAll() ([]entity.Location, error)
}

type LocationRepository struct {
	DB *database.Postgres
}

func NewLocationRepository(db *database.Postgres) ILocationRepository {
	return &LocationRepository{
		DB: db,
	}
}

func (r *LocationRepository) Create(l *request.LocationRequest) error {
	if err := r.DB.GDB.Table("locations").Create(l).Error; err != nil {
		return err
	}

	return nil
}

func (r *LocationRepository) FindAll() ([]entity.Location, error) {
	var locations []entity.Location
	if err := r.DB.GDB.Table("locations").Find(&locations).Error; err != nil {
		return nil, err
	}

	return locations, nil
}
