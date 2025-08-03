package internal

import (
	"fmt"

	"github.com/Nikitha2404/go-grpc-api/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

func NewDB() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

func (d *DB) AddOrder(order *orders.Order) error {
	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			return fmt.Errorf("duplicate order id present. id=%d", o.OrderId)
		}
	}
	d.collection = append(d.collection, order)
	return nil
}
