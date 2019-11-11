package router

import (
	"io"
	"os"
	"strings"

	. "backend-kendo-tutorial/apis"
	"backend-kendo-tutorial/middleware/token"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRoute() *gin.Engine {

	// 紀錄Log
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 設定Gin Mode (Local=>debug、測試機=>test、正式機=>production)
	gin.SetMode(viper.GetString("mode"))
	router := gin.Default()

	// 跨網域設定
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost"}
	config.AllowHeaders = []string{"auth_token", "content-type", "Access-Control-Allow-Origin"}
	config.AllowOrigins = strings.Split(viper.GetString("cors"), ",")

	router.Use(cors.New(config))

	// 提供container健康檢測用
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	apiv1 := router.Group("/api/v1")

	// 登入API
	apiAuth := apiv1.Group("/auth")
	apiAuth.POST("/login", Login)
	apiAuth.POST("/refresh", RefreshToken)

	apiv1.Use(token.TokenAuthMiddleware())
	{
		apiAuth := apiv1.Group("/auth")
		apiAuth.GET("/user", GetAuthUser)

	}

	return router
}
