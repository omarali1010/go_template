package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/omaraliali1010/go_template/domain"
	"github.com/omaraliali1010/go_template/internal/db"
)

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(dbtx db.DBTX) domain.UserRepository {
	return &userRepository{
		queries: db.New(dbtx),
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := ur.queries.CreateUser(ctx, db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	return err
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
func (ur *userRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	u, err := ur.queries.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
