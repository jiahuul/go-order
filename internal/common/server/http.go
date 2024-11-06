package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHTTPServer(serviceName string, wrapper func(router *gin.Engine)) {
	addr := viper.Sub(serviceName).GetString("http-addr")
	if addr == "" {
		// todo: add warning
	}
	RunHTTPServerOnAddr(addr, wrapper)
}

func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRoute := gin.New()
	wrapper(apiRoute)
	apiRoute.Group("/api")
	apiRoute.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	err := apiRoute.Run(addr)
	if err != nil {
		panic(err)
	}
}
