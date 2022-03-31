package postgres

import (
	"errors"
	"gorm.io/gorm"
	"mini-rest-api/domain"
	"mini-rest-api/repository"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return userRepository{
		db,
	}
}

func (u userRepository) FindUserByID(id int) (*domain.User, error) {
	var user domain.User
	err := u.db.First(&user, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) CheckLogin(username string, password string) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
