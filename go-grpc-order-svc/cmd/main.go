package main

import (
	"fmt"
	"log"
	"net"

	"github.com/theninjacoder-uz/go-grpc-order-svc/pkg/client"
	"github.com/theninjacoder-uz/go-grpc-order-svc/pkg/config"
	"github.com/theninjacoder-uz/go-grpc-order-svc/pkg/db"
	"github.com/theninjacoder-uz/go-grpc-order-svc/pkg/pb"
	"github.com/theninjacoder-uz/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	fmt.Println("Order svc on:", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpServer, &s)

	if err := grpServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
