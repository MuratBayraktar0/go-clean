package repository

import (
	"context"
	"golang-clean/domain"
)

type Database interface {
	Get(ctx context.Context, id string) (*domain.UserEntity, error)
	GetAll(ctx context.Context) (*[]domain.UserEntity, error)
	Create(ctx context.Context, user domain.UserEntity) error
	Update(ctx context.Context, user domain.UserEntity) error
}
