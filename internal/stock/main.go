package main

import (
	"github.com/jiahuul/go-order/common/server"
	"github.com/jiahuul/go-order/genproto/stockpb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("stock.service-name")
	server.RunGRPCServer(serviceName, func (service *grpc.Server)) {
		stockpb.RegisterStockServiceServer(service, )
	}
}
