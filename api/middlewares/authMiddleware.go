package middlewares

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"mini-rest-api/common/exception"
	"mini-rest-api/common/helper"
	external "mini-rest-api/infra"
	"mini-rest-api/repository/postgres"
	"strconv"
)

func AuthCheck(token string, c echo.Context) (valid bool, err error) {
	h := helper.NewHelperFunction()

	valid = true

	tokenClaim, errParse := h.ParseJwt(token)
	if errParse != nil {

		return false, errParse
	}
	if tokenClaim.UserID == "" {
		return false, errors.New("invalid user")
	}

	newDB := external.NewGormDB()
	userRepo := postgres.NewUserRepository(newDB)

	userID, err := strconv.Atoi(tokenClaim.UserID)
	if err != nil {
		return false, err
	}

	authID, err := uuid.Parse(tokenClaim.AuthID)
	if err != nil {
		return false, err
	}

	user, errGetUser := userRepo.FindUserByIDAndAuthID(userID, authID)
	if errGetUser != nil {
		exception.PanicIfNeeded(err)
	}
	if user != nil {
		c.Set("user_id", user.ID)
		c.Set("username", user.Username)
		c.Set("auth_id", user.AuthUUID)
		return
	}

	return
}

func RefreshCheck(token string, c echo.Context) (valid bool, err error) {
	h := helper.NewHelperFunction()

	valid = true

	tokenClaim, errParse := h.ParseRefreshJwt(token)
	if errParse != nil {

		return false, errParse
	}
	if tokenClaim.UserID == "" {
		return false, errors.New("invalid user")
	}

	newDB := external.NewGormDB()
	userRepo := postgres.NewUserRepository(newDB)

	userID, err := strconv.Atoi(tokenClaim.UserID)
	if err != nil {
		return false, err
	}

	authID, err := uuid.Parse(tokenClaim.AuthID)
	if err != nil {
		return false, err
	}

	user, errGetUser := userRepo.FindUserByIDAndAuthID(userID, authID)
	if errGetUser != nil {
		exception.PanicIfNeeded(err)
	}
	if user != nil {
		c.Set("user_id", user.ID)
		c.Set("username", user.Username)
		c.Set("auth_id", user.AuthUUID)
		return
	}

	return
}
