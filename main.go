package main

import (
	"fmt"
	"log"

	"github.com/Nikitha2404/go-grpc-api/protogen/golang/orders"
	"github.com/Nikitha2404/go-grpc-api/protogen/golang/product"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	orderItem := orders.Order{
		OrderId:    1,
		CustomerId: 1,
		IsActive:   true,
		OrderDate: &date.Date{
			Year: 2025, Month: 8, Day: 3,
		},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "Fanta", ProductType: product.ProductType_DRINK},
		},
	}

	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatalf("deserialisation error:%v", err)
	}

	fmt.Println(string(bytes))
}
