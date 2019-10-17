package token

import (
	"backend-kendo-tutorial/models/resp"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 取token
		authToken := c.Request.Header.Get("auth_token")

		// 如果token找不到就給錯
		if authToken == "" {
			respondWithError(401, "找不到token", c)
			return
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
