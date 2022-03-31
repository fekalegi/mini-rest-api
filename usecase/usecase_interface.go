package usecase

import (
	"github.com/google/uuid"
	"mini-rest-api/common/dto"
)

type UserService interface {
	CheckLogin(username string, password string) (dto.JsonResponses, error)
	AddUser(request dto.RequestAddUser) (dto.JsonResponses, error)
	RefreshToken(userID int, authID uuid.UUID) (dto.JsonResponses, error)
	DeleteAuth(userID int, authID uuid.UUID) (dto.JsonResponses, error)
}
