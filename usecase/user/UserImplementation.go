package user

import (
	"mini-rest-api/common/interfaces"
	"mini-rest-api/repository"
	"mini-rest-api/usecase"
)

type userImplementation struct {
	repo   repository.UserRepository
	helper interfaces.HelperInterface
}

func NewUserImplementation(repo repository.UserRepository, helper interfaces.HelperInterface) usecase.UserService {
	return &userImplementation{
		repo:   repo,
		helper: helper,
	}
}
