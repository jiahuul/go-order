package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jiahuul/go-order/common/config"
	"github.com/jiahuul/go-order/common/server"
	"github.com/jiahuul/go-order/ports"
	"github.com/spf13/viper"
	"log"
)

// 在main函数执行前自动执行，且不能被其他函数调用
func init() {
	err := config.NewViperConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Printf("%v", viper.Get("order"))
	serviceName := viper.GetString("order.service-name")

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
