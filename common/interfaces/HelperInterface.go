package interfaces

import "mini-rest-api/common/models"

type HelperInterface interface {
	CreateJwtTokenLogin(userID, username string) (token string, err error)
	CreateRefreshJwtTokenLogin(userID string) (token string, err error)
	ParseJwt(token string) (claims models.TokenClaims, err error)
}
