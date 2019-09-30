package service

import (
	"comuniapp/models"
)

type UserServices interface {
	CreateUser(user models.User) (int64, error)
	GetAllUser() ([]models.User, error)
	GetUserById(id int64) (*models.User, error)
	UpdateUser(user models.User) (bool, error)
	DeleteUser(user models.User) (bool, error)
}

type userStruct struct{}

func CreateUserServices() UserServices {
	return userStruct{}
}

func (userStruct) CreateUser(user models.User) (int64, error) {
	err := user.Insert()
	if err != nil {
		return 0, nil
	}
	return 1, nil
}

func (userStruct) GetAllUser() ([]models.User, error) {
	return models.GetAllUsers()
}

func (userStruct) GetUserById(id int64) (*models.User, error) {
	return models.GetUserById(id)
}

func (userStruct) UpdateUser(user models.User) (bool, error) {
	err := user.Update()
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (userStruct) DeleteUser(user models.User) (bool, error) {
	err := user.Delete()
	if err != nil {
		return false, nil
	}
	return true, nil
}
