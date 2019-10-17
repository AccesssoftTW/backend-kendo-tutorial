package services

import (
	"errors"

	"backend-kendo-tutorial/models/token"
	"backend-kendo-tutorial/models/user"

	"github.com/spf13/viper"
)

type AuthService struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (this AuthService) Login() (tokenModel token.Token, err error) {

	// admin檢查
	if this.Account == viper.GetString("admin.account") && this.Password == viper.GetString("admin.password") {
		var tokenService TokenService
		tokenModel.AuthToken, err = tokenService.GenerateToken("1")
		tokenModel.RefreshToken, err = tokenService.GenerateRefreshToken("1")
		return
	} else {
		err = errors.New("帳號碼密不正確")
	}
	return
}

func (this AuthService) GetAuthUser(tokenStr string) (authUser user.User, err error) {

	authUser.Name = "elsa"

	return
}
