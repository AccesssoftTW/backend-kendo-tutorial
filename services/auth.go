package services

import (
	"errors"

	"github.com/spf13/viper"
)

type AuthService struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (this AuthService) Login() (token TokenService, err error) {

	// admin檢查
	if this.Account == viper.GetString("admin.account") && this.Password == viper.GetString("admin.password") {
		return
	} else {
		err = errors.New("帳號碼密不正確")
	}
	return
}
