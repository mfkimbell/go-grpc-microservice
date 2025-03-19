package service

import (
	"context"
	"database/sql"

	"github.com/mfkimbell/go-grpc-microservice/services/common/genproto/orders"
)

type DatabaseOrderService struct {
	db *sql.DB
}

func (d *DatabaseOrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	_, err := d.db.Exec("INSERT INTO orders (customer_id, product_id, quantity) VALUES (?, ?, ?)", order.CustomerID, order.ProductID, order.Quantity)
	return err
}

func (d *DatabaseOrderService) GetOrders(ctx context.Context) []*orders.Order {
	rows, _ := d.db.Query("SELECT order_id, customer_id, product_id, quantity FROM orders")
	defer rows.Close()

	var orderList []*orders.Order 
	for rows.Next() {
		order := &orders.Order{}
		rows.Scan(&order.OrderID, &order.CustomerID, &order.ProductID, &order.Quantity)
		orderList = append(orderList, order)
	}
	return orderList
}
