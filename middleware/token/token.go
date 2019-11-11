package token

import (
	"backend-kendo-tutorial/databases"
	"backend-kendo-tutorial/models/resp"
	"backend-kendo-tutorial/models/user"
	"backend-kendo-tutorial/services"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
			claims, err := tokenService.GetTokenInfo(authToken)
			if err != nil {

				// token不合法，報錯
				respondWithError(401, "此帳號憑證不合法或已過期", c)
				return
			}

			// 檢查這個用戶是否過期了
			var userEntity user.User
			var userId uint64
			userId, err = strconv.ParseUint(claims["userId"].(string), 10, 64)
			if err != nil {
				respondWithError(401, err.Error(), c)
				return
			}
			userEntity.ID = uint(userId)
			if err = databases.Eloquent.First(&userEntity, userEntity.ID).Error; gorm.IsRecordNotFoundError(err) {
				err = errors.New("此用戶不存在")
				respondWithError(401, err.Error(), c)
				return
			}
			if err != nil {
				respondWithError(401, err.Error(), c)
				return
			}
			if userEntity.ExpiredAt != nil {
				var t1 = userEntity.ExpiredAt.Unix()
				var t2 = time.Now().Unix()
				if t1 < t2 {
					respondWithError(401, "帳號已過期", c)
					return
				}
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
