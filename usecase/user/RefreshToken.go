package user

import (
	"github.com/google/uuid"
	"mini-rest-api/common/command"
	"mini-rest-api/common/dto"
	"strconv"
)

func (u userImplementation) RefreshToken(userID int, authID uuid.UUID) (dto.JsonResponses, error) {
	user, err := u.repo.FindUserByIDAndAuthID(userID, authID)
	if user == nil && err == nil {
		return command.NotFoundResponses("User not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), nil
	}

	newAuthID := uuid.New()
	err = u.repo.UpdateAuthUUID(user.ID, newAuthID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), nil
	}
	conv := strconv.Itoa(user.ID)
	token, errCreateToken := u.helper.CreateJwtTokenLogin(conv, user.Username, newAuthID)
	if errCreateToken != nil {
		return command.InternalServerResponses(errCreateToken.Error()), err
	}

	response := dto.ResponseRefreshToken{
		AccessToken: token,
	}
	return command.SuccessResponses(response), nil
}
