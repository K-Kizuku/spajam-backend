package repository

import (
	"context"

	"github.com/K-Kizuku/spajam-backend/db/sql/query"
	"github.com/K-Kizuku/spajam-backend/internal/domain/entity"
	"github.com/K-Kizuku/spajam-backend/internal/domain/repository"
	"github.com/K-Kizuku/spajam-backend/pkg/errors"
)

type UserRepository struct {
	queries *query.Queries
}

func NewUserRepository(queries *query.Queries) repository.IUserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.HandleDBError(err)
	}
	e := &entity.User{
		ID:       user.UserID,
		Username: user.Name,
		Password: user.HashedPassword,
		Email:    user.Mail,
		Code:     user.Code,
	}
	return e, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.HandleDBError(err)
	}
	e := &entity.User{
		ID:       user.UserID,
		Username: user.Name,
		Password: user.HashedPassword,
		Email:    user.Mail,
	}
	return e, nil
}

func (r *UserRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	u, err := r.queries.CreateUser(ctx, query.CreateUserParams{
		UserID:         user.ID,
		Mail:           user.Email,
		Name:           user.Username,
		HashedPassword: user.Password,
	})
	if err != nil {
		return nil, errors.HandleDBError(err)
	}
	createdUser := &entity.User{
		ID:       u.UserID,
		Username: u.Name,
		Password: u.HashedPassword,
		Email:    u.Mail,
	}
	return createdUser, nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id, password string) error {
	err := r.queries.UpdatePassword(ctx, query.UpdatePasswordParams{
		UserID:         id,
		HashedPassword: password,
	})
	if err != nil {
		return errors.HandleDBError(err)
	}
	return nil
}

func (r *UserRepository) UpdateCode(ctx context.Context, id, code string) error {
	err := r.queries.UpdateCode(ctx, query.UpdateCodeParams{
		UserID: id,
		Code:   code,
	})
	if err != nil {
		return errors.HandleDBError(err)
	}
	return nil
}
