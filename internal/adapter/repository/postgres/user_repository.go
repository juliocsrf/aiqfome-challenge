package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/database"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
)

type UserRepositoryImpl struct {
	Queries *database.Queries
}

func NewUserRepository(queries *database.Queries) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Queries: queries,
	}
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	ctx := context.Background()
	user, err := u.Queries.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	userEntity, err := entity.NewUserWithId(user.ID.String(), user.Name, user.Email, user.Password)
	if err != nil {
		return nil, fmt.Errorf("error creating user entity: %s", err.Error())
	}

	return userEntity, nil
}

func (u *UserRepositoryImpl) FindByID(id string) (*entity.User, error) {
	ctx := context.Background()

	userUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %s", err.Error())
	}

	user, err := u.Queries.FindUserById(ctx, userUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	userEntity, err := entity.NewUserWithId(user.ID.String(), user.Name, user.Email, user.Password)
	if err != nil {
		return nil, fmt.Errorf("error creating user entity: %s", err.Error())
	}

	return userEntity, nil
}
