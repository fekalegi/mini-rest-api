package user

import (
	"mini-rest-api/common/command"
	"mini-rest-api/common/dto"
	"strconv"
)

func (u userImplementation) RefreshToken(userID int) (dto.JsonResponses, error) {
	user, err := u.repo.FindUserByID(userID)
	if user == nil && err == nil {
		return command.NotFoundResponses("User not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), nil
	}

	conv := strconv.Itoa(user.ID)
	token, errCreateToken := u.helper.CreateJwtTokenLogin(conv, user.Username)
	if errCreateToken != nil {
		return command.InternalServerResponses(errCreateToken.Error()), err
	}

	response := dto.ResponseRefreshToken{
		AccessToken: token,
	}
	return command.SuccessResponses(response), nil
}
