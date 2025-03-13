package usecase

import (
	"fmt"
	"indico-technical-test/app/entity"
	"indico-technical-test/app/repository"

	"gorm.io/gorm"
)

type IUserUsecase interface {
	FindAll() ([]entity.User, error)
	FindMe(token string) (*entity.User, error)
}

type UserUsecase struct {
	UserRepository repository.IUserRepository
}

func NewUserUsecase(userRepo repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		UserRepository: userRepo,
	}
}

func (u *UserUsecase) FindAll() ([]entity.User, error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %v", err)
	}

	return users, nil
}

func (u *UserUsecase) FindMe(token string) (*entity.User, error) {
	user, err := u.UserRepository.FindMe(token)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entity.User{}, nil
		}
		return nil, fmt.Errorf("error fetching user: %v", err)
	}

	return user, nil
}
