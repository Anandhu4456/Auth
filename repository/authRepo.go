package repository

import (
	"auth/domain"
	"auth/dto"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, req *dto.UserCreateRequest, hashedPassword string) (*domain.User, error)
}

type AuthRepoImpl struct {
	dbPool *pgxpool.Pool
}

// repo constructor
func NewAuthRepo(dbPool *pgxpool.Pool) AuthRepo {
	return &AuthRepoImpl{
		dbPool: dbPool,
	}
}

func (r *AuthRepoImpl) CreateUser(ctx context.Context, req *dto.UserCreateRequest, hashedPassword string) (*domain.User, error) {

	var id int64

	query := `
			INSERT INTO users(name, username, email, password, phone_number, address) VALUES($1, $2, $3, $4, $5, $6) RETURNING id
	`

	if err := r.dbPool.QueryRow(ctx, query, req.Name, req.Username, req.Email, hashedPassword, req.PhoneNumber, req.Address).Scan(&id); err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:          id,
		Name:        req.Name,
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		CreatedAt:   time.Now(),
	}
	return user, nil
}
