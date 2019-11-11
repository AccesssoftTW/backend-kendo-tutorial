package services

import (
	"errors"
	"strconv"

	"backend-kendo-tutorial/databases"
	"backend-kendo-tutorial/models/token"
	"backend-kendo-tutorial/models/user"

	"github.com/jinzhu/gorm"
)

type AuthService struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (this AuthService) Login() (tokenModel token.Token, err error) {

	// 檢查用戶
	var userEntity user.User
	if err = databases.Eloquent.Where(&user.User{Account: this.Account, Password: this.Password}).First(&userEntity).Error; gorm.IsRecordNotFoundError(err) {
		err = errors.New("此用戶不存在")
		return
	}
	if err != nil {
		return
	}

	var tokenService TokenService
	// 產生token
	tokenModel.AuthToken, err = tokenService.GenerateToken(strconv.Itoa(int(userEntity.ID)))
	if err != nil {
		err = errors.New("產生token時發生了錯誤")
		return
	}

	// 產生refresh token
	tokenModel.RefreshToken, err = tokenService.GenerateRefreshToken(strconv.Itoa(int(userEntity.ID)))
	if err != nil {
		err = errors.New("產生refresh token時發生了錯誤")
		return
	}

	return
}

func (this AuthService) GetAuthUser(tokenStr string) (authUser user.User, err error) {

	// 取得token中的資料
	var tokenService TokenService
	claims, err := tokenService.GetTokenInfo(tokenStr)
	if err != nil {
		err = errors.New("無法取得用戶資訊，憑證可能已過期。")
		return
	}

	// 取得用戶ID
	var userId uint64
	userId, err = strconv.ParseUint(claims["userId"].(string), 10, 64)
	if err != nil {
		return
	}

	// 撈用戶資料
	authUser.ID = uint(userId)
	if err = databases.Eloquent.First(&authUser, authUser.ID).Error; gorm.IsRecordNotFoundError(err) {
		err = errors.New("此憑證的用戶不存在")
		return
	}
	if err != nil {
		return
	}

	return
}

func (this AuthService) RefreshToken(refreshToken string) (token token.Token, err error) {

	var tokenService TokenService
	claims, err := tokenService.GetTokenInfo(refreshToken)
	if err == nil {
		userId, ok := claims["userId"].(string)
		if !ok {
			err = errors.New("此憑證不合法，沒有用戶資訊")
			return
		}

		// 重新產生token
		token.AuthToken, err = tokenService.GenerateToken(userId)
		if err != nil {
			err = errors.New("用戶的憑證產生失敗，請重新試一次")
			return
		}

		// 重新產生refresh token
		token.RefreshToken, err = tokenService.GenerateRefreshToken(userId)
		if err != nil {
			err = errors.New("用戶的[更新]憑證產生失敗，請重新試一次")
			return
		}

		return
	} else {
		err = errors.New("憑證不合法")
		return
	}
}
