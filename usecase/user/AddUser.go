package user

import (
	"mini-rest-api/common/command"
	"mini-rest-api/common/dto"
)

func (u userImplementation) AddUser(request dto.RequestAddUser) (dto.JsonResponses, error) {
	return command.SuccessResponses(""), nil
}
