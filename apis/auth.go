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
