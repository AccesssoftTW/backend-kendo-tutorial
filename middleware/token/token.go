package token

import (
	"backend-kendo-tutorial/models/resp"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()
	}
}

func respondWithError(code int, message string, c *gin.Context) {

	var resp resp.ApiResponse
	resp.Result = message

	c.JSON(code, resp)

	c.Abort()
}
