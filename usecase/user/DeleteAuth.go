package user

import (
	"github.com/google/uuid"
	"mini-rest-api/common/command"
	"mini-rest-api/common/dto"
)

func (u userImplementation) DeleteAuth(userID int, authID uuid.UUID) (dto.JsonResponses, error) {
	user, err := u.repo.FindUserByIDAndAuthID(userID, authID)
	if user == nil && err == nil {
		return command.NotFoundResponses("User not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), nil
	}

	nilUUID := uuid.Nil
	err = u.repo.UpdateAuthUUID(userID, nilUUID)

	return command.SuccessResponses("success"), nil
}
