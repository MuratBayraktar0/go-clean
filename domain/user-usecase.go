package domain

import "context"

type UserUsecase interface {
	Get(ctx context.Context, id string) (*UserDTO, error)
	GetAll(ctx context.Context) (*[]UserDTO, error)
	Create(ctx context.Context, user UserDTO) error
	Update(ctx context.Context, id string, user UserDTO) error
}
