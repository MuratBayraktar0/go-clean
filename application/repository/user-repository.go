package repository

import (
	"context"
	"golang-clean/domain"
)

type userRepository struct {
	db             Database
	contextTimeout context.Context
}

func NewRepository(timeout context.Context, database Database) *userRepository {
	return &userRepository{
		db:             database,
		contextTimeout: timeout,
	}
}

func (u *userRepository) Get(ctx context.Context, id string) (*domain.UserEntity, error) {
	userEntity, err := u.db.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return userEntity, nil
}

func (u *userRepository) GetAll(ctx context.Context) (*[]domain.UserEntity, error) {
	usersEntity, err := u.db.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return usersEntity, nil
}

func (u *userRepository) Create(ctx context.Context, user domain.UserEntity) error {
	err := u.db.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Update(ctx context.Context, user domain.UserEntity) error {
	err := u.db.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
