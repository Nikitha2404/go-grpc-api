package internal

import (
	"context"
	"log"

	"github.com/Nikitha2404/go-grpc-api/protogen/golang/orders"
)

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func NewOrderService(db *DB) OrderService {
	return OrderService{
		db: db,
	}
}

func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Println("Recieved an add-order request")

	err := o.db.AddOrder(req.GetOrder())
	return &orders.Empty{}, err
}
