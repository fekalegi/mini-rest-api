package usecase

import "mini-rest-api/common/dto"

type UserService interface {
	CheckLogin(username string, password string) (dto.JsonResponses, error)
	AddUser(request dto.RequestAddUser) (dto.JsonResponses, error)
	RefreshToken(userID int) (dto.JsonResponses, error)
}
