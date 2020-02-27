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

func (this UserService) UpdateUser(userEntity *user.User) (err error) {

	if err = databases.Eloquent.Model(&userEntity).Updates(userEntity).Error; err != nil {
		err = errors.New("更新發生問題，請找 winston")
		return
	}

	return
}

func (this UserService) DeleteUser(id string) (err error) {

	if err = databases.Eloquent.Where("id = ?", id).Delete(&user.User{}).Error; err != nil {
		err = errors.New("刪除發生問題，請找 winston")
		return
	}

	return
}
