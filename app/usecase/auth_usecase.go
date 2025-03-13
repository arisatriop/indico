package usecase

import (
	"fmt"
	"indico-technical-test/app/entity"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/app/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type IAuthUsecase interface {
	Register(*request.RegisterRequest) error
	Login(*request.LoginRequest) (string, error)
}

type AuthUsecase struct {
	AuthRepository repository.IAuthRepository
}

func NewAuthUsecase(authRepo repository.IAuthRepository) IAuthUsecase {
	return &AuthUsecase{
		AuthRepository: authRepo,
	}
}

func (u *AuthUsecase) Register(req *request.RegisterRequest) error {
	user := entity.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Roles:    req.Roles,
	}

	err := u.AuthRepository.Register(&user)
	if err != nil {
		return fmt.Errorf("error registering user: %v", err)
	}

	return nil
}

func (u *AuthUsecase) Login(req *request.LoginRequest) (string, error) {

	user, err := u.AuthRepository.Login(req)
	if err != nil {
		return "", fmt.Errorf("error logging in: %v", err)
	}

	exp := time.Now().Add(1 * time.Hour)
	claims := &entity.Claims{
		Username: req.Username,
		Role:     user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	user.Token = tokenStr
	if err := u.AuthRepository.Update(user); err != nil {
		return "", fmt.Errorf("error updating user: %v", err)
	}

	return tokenStr, nil
}
