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

func (d *DB) GetOrder(orderId uint64) *orders.Order {
	for _, o := range d.collection {
		if o.OrderId == orderId {
			return o
		}
	}
	return nil
}

func (d *DB) DeleteOrder(orderId uint64) {
	filtered := make([]*orders.Order, len(d.collection)-1)
	for i := range d.collection {
		if d.collection[i].OrderId != orderId {
			filtered = append(filtered, d.collection[i])
		}
	}
	d.collection = filtered
}

func (d *DB) UpdateOrder(order *orders.Order) {
	for i, o := range d.collection {
		if o.OrderId == order.OrderId {
			d.collection[i] = order
			return
		}
	}
}
