package service

import (
	"auth/dto"
	"auth/repository"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateUser(ctx context.Context, req *dto.UserCreateRequest) (*dto.UserCreateResponse, error)
}

type AuthServiceImpl struct {
	authRepo repository.AuthRepo
}

// authservice constructor
func NewAuthService(authRepo repository.AuthRepo) AuthService {
	return &AuthServiceImpl{
		authRepo: authRepo,
	}
}

func (s *AuthServiceImpl) CreateUser(ctx context.Context, req *dto.UserCreateRequest) (*dto.UserCreateResponse, error) {

	//password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password hashing failed : %w", err)
	}

	user, err := s.authRepo.CreateUser(ctx, req, string(hashedPassword))
	if err != nil {
		return nil, fmt.Errorf("user creation failed : %w", err)
	}

	response := &dto.UserCreateResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.Email,
		Address:     user.Address,
	}

	return response, err
}
