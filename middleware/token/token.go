package token

import (
	"backend-kendo-tutorial/models/resp"
	"backend-kendo-tutorial/services"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 取token
		authToken := c.Request.Header.Get("auth_token")
		if authToken == "" {

			// 沒有token，報錯
			respondWithError(401, "找不到憑證", c)
			return
		} else {

			// 取得token的話就驗證
			var tokenService services.TokenService
			_, err := tokenService.GetTokenInfo(authToken)
			if err != nil {

				// token不合法，報錯
				respondWithError(401, "此帳號憑證不合法或已過期", c)
				return
			}
		}

		c.Next()
	}
}

func respondWithError(code int, message string, c *gin.Context) {

	var resp resp.ApiResponse
	resp.Result = message

	c.JSON(code, resp)

	c.Abort()
}
