package services

import (
	"backend-kendo-tutorial/databases"
	"backend-kendo-tutorial/models/user"
	"errors"
)

type UserService struct {
}

func (this UserService) GetUser() (userEntities []user.User, err error) {

	if err = databases.Eloquent.Find(&userEntities).Error; err != nil {
		err = errors.New("查詢發生問題，請找 winston")
		return
	}

	return
}

func (this UserService) AddUser(userEntity *user.User) (err error) {

	if err = databases.Eloquent.Create(&userEntity).Error; err != nil {
		err = errors.New("新增發生問題，請找 winston")
		return
	}

	return
}
