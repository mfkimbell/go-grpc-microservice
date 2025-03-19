package main

import (
	"log"
	"net"

	handler "github.com/mfkimbell/go-grpc-microservice/services/orders/handler/orders"
	"github.com/mfkimbell/go-grpc-microservice/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// ### If we were using a real database
	// db := sql.Open("postgres", "my-database-connection-string")
	// ordersService := &service.DatabaseOrderService{db: db} // ✅ Uses PostgreSQL
	// NewGrpcOrdersService(grpcServer, ordersService)

	// Handlers only depend on OrderService, not a specific implementation.

	// register our grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}