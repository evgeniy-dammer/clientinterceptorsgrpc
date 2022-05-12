package main

import (
	"fmt"
	"net"

	"github.com/evgeniy-dammer/clientinterceptorsgrpc/productservice/handlers"
	productservice "github.com/evgeniy-dammer/clientinterceptorsgrpc/productservice/proto"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":1111")

	if err != nil {
		fmt.Println(err)
	}

	defer listen.Close()

	productServ := handlers.ProductServiceServer{}

	grpcServer := grpc.NewServer()

	productservice.RegisterProductServiceServer(grpcServer, &productServ)

	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
