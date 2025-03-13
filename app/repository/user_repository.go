package repository

import (
	"indico-technical-test/app/entity"
	"indico-technical-test/database"
)

type IUserRepository interface {
	FindAll() ([]entity.User, error)
	FindMe(token string) (*entity.User, error)
}

type UserRepository struct {
	DB *database.Postgres
}

func NewUserRepository(db *database.Postgres) IUserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.DB.GDB.Table("users").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindMe(token string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.GDB.Table("users").Where("token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
