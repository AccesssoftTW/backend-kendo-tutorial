package apis

import (
	"backend-kendo-tutorial/models/resp"
	"backend-kendo-tutorial/services"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	// 取出Body資料
	decoder := json.NewDecoder(c.Request.Body)
	var authService services.AuthService
	err := decoder.Decode(&authService)

	// 登入驗證
	token, err := authService.Login()
	if err != nil {

		var resp resp.ApiResponse
		resp.Result = err.Error()

		c.JSON(400, resp)
		return
	}

	c.JSON(200, token)
}

func GetAuthUser(c *gin.Context) {

	// 取token
	authToken := c.Request.Header.Get("auth_token")

	// 取得token裡面的user
	var authService services.AuthService
	authUser, err := authService.GetAuthUser(authToken)
	if err != nil {

		var resp resp.ApiResponse
		resp.Result = err.Error()

		c.JSON(400, resp)
		return
	}
	c.JSON(200, authUser)
}

func RefreshToken(c *gin.Context) {

	// 取refresh_token
	refreshToken := c.Request.Header.Get("refresh_token")
	if refreshToken == "" {

		var resp resp.ApiResponse
		resp.Result = "取得重整憑證失敗"
		c.JSON(400, resp)
		return
	}
	// 驗證
	var tokenService services.TokenService
	ok, err := tokenService.ValidateToken(refreshToken)
	if err != nil {

		var resp resp.ApiResponse
		resp.Result = err.Error()
		c.JSON(400, resp)
		return
	}

	if ok {
		var authService services.AuthService
		token, err := authService.RefreshToken(refreshToken)
		if err != nil {

			var resp resp.ApiResponse
			resp.Result = err.Error()
			c.JSON(400, resp)
			return
		}

		c.JSON(200, token)
		return
	} else {

		var resp resp.ApiResponse
		resp.Result = err.Error()
		c.JSON(400, resp)
		return
	}
}
