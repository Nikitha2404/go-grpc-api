package internal

import (
	"context"
	"fmt"
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
	log.Println("Exisiting Data: ", o.db.collection)
	log.Println("Recieved an add-order request")

	err := o.db.AddOrder(req.GetOrder())
	log.Println("Exisiting Data: ", o.db.collection)
	return &orders.Empty{}, err
}

func (o *OrderService) DeleteOrder(_ context.Context, req *orders.PayloadWithOrderId) (*orders.Empty, error) {
	log.Println("Exisiting Data: ", o.db.collection)
	log.Println("Recieved a delete-order request")

	o.db.DeleteOrder(req.OrderId)
	log.Println("Exisiting Data: ", o.db.collection)
	return &orders.Empty{}, nil
}

func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderId) (*orders.PayloadWithSingleOrder, error) {
	log.Println("Exisiting Data: ", o.db.collection)
	log.Println("Received get-order request")

	orderDetails := o.db.GetOrder(req.OrderId)
	if orderDetails == nil {
		return nil, fmt.Errorf("order details not found for id=%d", req.OrderId)
	}
	return &orders.PayloadWithSingleOrder{
		Order: orderDetails,
	}, nil
}

func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Println("Exisiting Data: ", o.db.collection)
	log.Println("Recieved update-order request")

	o.db.UpdateOrder(req.GetOrder())
	log.Println("Exisiting Data: ", o.db.collection)
	return &orders.Empty{}, nil
}
