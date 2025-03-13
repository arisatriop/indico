package repository

import (
	"indico-technical-test/app/entity"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/database"
)

type IAuthRepository interface {
	Register(*entity.User) error
	Login(*request.LoginRequest) (*entity.User, error)
	Update(*entity.User) error
}

type AuthRepository struct {
	DB *database.Postgres
}

func NewAuthRepository(db *database.Postgres) IAuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) Register(u *entity.User) error {
	if err := r.DB.GDB.Table("users").Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) Login(req *request.LoginRequest) (*entity.User, error) {
	var user entity.User
	if err := r.DB.GDB.Table("users").Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) Update(u *entity.User) error {
	if err := r.DB.GDB.Table("users").Save(u).Error; err != nil {
		return err
	}

	return nil
}
