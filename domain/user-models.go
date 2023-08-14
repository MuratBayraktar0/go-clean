package domain

import "time"

type UserDTO struct {
	Name  string
	Email string
}

type UserEntity struct {
	ID        string
	Name      string
	Email     string
	UpdatedAt time.Time
	CreatedAt time.Time
}
