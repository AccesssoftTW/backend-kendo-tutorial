package apis

import (
	"backend-kendo-tutorial/models/resp"
	"backend-kendo-tutorial/services"
	"encoding/json"

	"backend-kendo-tutorial/models/user"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	var authService services.UserService
	users, err := authService.GetUser()
	if err != nil {

		var resp resp.ApiResponse
		resp.Result = err.Error()

		c.JSON(400, resp)
		return
	}

	c.JSON(200, users)
}

func AddUser(c *gin.Context) {

	// 取出Body資料
	decoder := json.NewDecoder(c.Request.Body)
	var userEntity user.User
	err := decoder.Decode(&userEntity)

	var authService services.UserService
	err = authService.AddUser(&userEntity)
	if err != nil {

		var resp resp.ApiResponse
		resp.Result = err.Error()

		c.JSON(400, resp)
		return
	}

	c.JSON(200, userEntity)
}

func UpdateUser(c *gin.Context) {

	// 取出Body資料
	decoder := json.NewDecoder(c.Request.Body)
	var userEntity user.User
	err := decoder.Decode(&userEntity)

	var authService services.UserService
	err = authService.UpdateUser(&userEntity)
	if err != nil {

		var resp resp.ApiResponse
		resp.Result = err.Error()

		c.JSON(400, resp)
		return
	}

	c.JSON(200, userEntity)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	var authService services.UserService
	err := authService.DeleteUser(id)
	if err != nil {

		var resp resp.ApiResponse
		resp.Result = err.Error()

		c.JSON(400, resp)
		return
	}

	c.JSON(200, nil)
}
