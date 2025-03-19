package types

import (
	"context"

	"github.com/mfkimbell/go-grpc-microservice/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
