package main

import (
	"context"
	"fmt"

	interceptors "github.com/evgeniy-dammer/clientinterceptorsgrpc/client/interceptors"

	productservice "github.com/evgeniy-dammer/clientinterceptorsgrpc/productservice/proto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial(
		"localhost:1111",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				interceptors.DateLogClientInterceptor,
				interceptors.MethodLogClientInterceptor,
			),
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	defer connection.Close()

	productServ := productservice.NewProductServiceClient(connection)

	response1, err1 := productServ.FindAll(context.Background(), &productservice.FindAllRequest{})

	if err1 != nil {
		fmt.Println(err)
	} else {
		products := response1.Products

		fmt.Println("Product List")

		for _, product := range products {
			fmt.Println("Id: ", product.Id)
			fmt.Println("Name: ", product.Name)
			fmt.Println("Price: ", product.Price)
			fmt.Println("Quantity: ", product.Quantity)
			fmt.Println("Status: ", product.Status)
			fmt.Println("========================")
		}
	}

	response2, err2 := productServ.Search(context.Background(), &productservice.SearchRequest{Keyword: "vi"})

	if err2 != nil {
		fmt.Println(err)
	} else {
		products := response2.Products

		fmt.Println("Search Result")

		for _, product := range products {
			fmt.Println("Id: ", product.Id)
			fmt.Println("Name: ", product.Name)
			fmt.Println("Price: ", product.Price)
			fmt.Println("Quantity: ", product.Quantity)
			fmt.Println("Status: ", product.Status)
			fmt.Println("========================")
		}
	}
}
