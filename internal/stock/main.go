package main

import (
	"context"
	"github.com/jiahuul/go-order/common/config"
	"github.com/jiahuul/go-order/common/genproto/stockpb"
	"github.com/jiahuul/go-order/common/server"
	"github.com/jiahuul/go-order/ports"
	"github.com/jiahuul/go-order/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// 要初始化viper之后，viper才能正常工作哦
func init() {
	err := config.NewViperConfig()
	if err != nil {
		logrus.Panic(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serviceType := viper.GetString("stock.service-to-run")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := service.NewApplication(ctx)
	logrus.Println(serviceType)
	if serviceType == "grpc" {
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			stockpb.RegisterStockServiceServer(server, ports.NewGRPCServer(app))
		})
	} else {
		logrus.Panic("unexpected server type")
	}

}
