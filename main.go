package main

import (
	"fmt"

	"backend-kendo-tutorial/router"

	"github.com/spf13/viper"
)

func main() {

	LoadConfig()

	// 啟動Gin
	app := router.InitRoute()
	app.Run(viper.GetString("server.port"))
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
