package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Nikitha2404/go-grpc-api/protogen/golang/orders"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	orderSvcAddr := "localhost:50051"
	conn, err := grpc.NewClient(orderSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service:%v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	if err = orders.RegisterOrdersHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to registeer order service:%v", err)
	}

	addr := "localhost:8080"
	fmt.Println("API gateway server running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("gateway server closed abruptly:%v", err)
	}
}
