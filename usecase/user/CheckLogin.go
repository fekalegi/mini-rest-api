package user

import (
	"mini-rest-api/common/command"
	"mini-rest-api/common/dto"
	"strconv"
)

func (u userImplementation) CheckLogin(username string, password string) (dto.JsonResponses, error) {
	user, err := u.repo.CheckLogin(username, password)
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

	refreshToken, errCreateToken := u.helper.CreateRefreshJwtTokenLogin(conv)
	if errCreateToken != nil {
		return command.InternalServerResponses(errCreateToken.Error()), err
	}
	response := dto.ResponseLogin{
		RefreshToken: refreshToken,
		AccessToken:  token,
	}
	return command.SuccessResponses(response), nil
}
