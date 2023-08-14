package domain

import "context"

type UserRepository interface {
	Get(ctx context.Context, id string) (*UserEntity, error)
	GetAll(ctx context.Context) (*[]UserEntity, error)
	Create(ctx context.Context, user UserEntity) error
	Update(ctx context.Context, user UserEntity) error
}
