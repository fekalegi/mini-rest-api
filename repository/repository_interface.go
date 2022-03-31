package repository

import (
	"mini-rest-api/domain"
)

type UserRepository interface {
	CheckLogin(username string, password string) (*domain.User, error)
	FindUserByID(id int) (*domain.User, error)
}
