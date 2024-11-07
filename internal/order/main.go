package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jiahuul/go-order/common/config"
	"github.com/jiahuul/go-order/common/genproto/orderpb"
	"github.com/jiahuul/go-order/common/server"
	"github.com/jiahuul/go-order/ports"
	"github.com/jiahuul/go-order/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

// 在main函数执行前自动执行，且不能被其他函数调用

func init() {
	err := config.NewViperConfig()
	if err != nil {
		logrus.Panic(err)
	}
}

func main() {
	log.Printf("%v", viper.Get("order"))
	serviceName := viper.GetString("order.service-name")
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	app := service.NewApplication(ctx)

	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		orderpb.RegisterOrderServiceServer(server, ports.NewGRPCServer(app))
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, ports.NewHTTPServer(), ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

}
