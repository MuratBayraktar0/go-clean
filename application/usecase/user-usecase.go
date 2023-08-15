package usecase

import (
	"context"
	"golang-clean/domain"
	"time"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout context.Context
}

func NewUserUsecase(timeout context.Context, userRepo domain.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Get(ctx context.Context, id string) (*domain.UserDTO, error) {
	userEntity, err := u.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	userDTO := ConvertEntitytoDTO(*userEntity)
	return &userDTO, nil
}

func (u *userUsecase) GetAll(ctx context.Context) (*[]domain.UserDTO, error) {
	usersEntity, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	usersDTO := []domain.UserDTO{}
	for _, userEntity := range *usersEntity {
		usersDTO = append(usersDTO, ConvertEntitytoDTO(userEntity))
	}

	return &usersDTO, nil
}

func (u *userUsecase) Create(ctx context.Context, user domain.UserDTO) error {
	userEntity := ConvertDTOtoEntity(user)
	userEntity.ID = uuid.New().String()
	userEntity.UpdatedAt = time.Now()
	userEntity.CreatedAt = time.Now()
	err := u.userRepo.Create(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Update(ctx context.Context, id string, user domain.UserDTO) error {
	userEntity := ConvertDTOtoEntity(user)
	userEntity.ID = id
	userEntity.UpdatedAt = time.Now()
	err := u.userRepo.Update(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}

func ConvertEntitytoDTO(userEntity domain.UserEntity) domain.UserDTO {
	return domain.UserDTO{
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
}

func ConvertDTOtoEntity(userDTO domain.UserDTO) domain.UserEntity {
	return domain.UserEntity{
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}
}
