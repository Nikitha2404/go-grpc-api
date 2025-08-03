package main

import (
	"log"
	"net"

	"github.com/Nikitha2404/go-grpc-api/internal"
	"github.com/Nikitha2404/go-grpc-api/protogen/golang/orders"
	"google.golang.org/grpc"
)

func main() {
	const addr = "localhost:50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	db := internal.NewDB()
	orderSvc := internal.NewOrderService(db)

	orders.RegisterOrdersServer(server, &orderSvc)

	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
