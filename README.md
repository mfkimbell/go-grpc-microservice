# go-grpc-microservice
```
brew install protobuf
```

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```
export PATH="$PATH:$(go env GOPATH)/bin"
```

```
go get google.golang.org/grpc
```

The handler layer processes the request and calls the service.
The service layer does NOT know anything about gRPC or HTTP.

```
services/orders/
├── handler/        # HTTP & gRPC request handling
│   ├── http.go
│   ├── grpc.go
├── service/        # Business logic
│   ├── orders.go
│   ├── payment.go  # (New: Handles payments for orders)
├── types/          # Shared data types
│   ├── types.go
```

```
services/orders/
├── handler/
│   ├── http.go
│   ├── grpc.go
├── service/
│   ├── orders.go
│   ├── fulfillment.go  # Handles order shipping
│   ├── payment.go      # Handles payments for orders
├── repository/        # (New: Stores orders in a database)
│   ├── postgres.go
│   ├── dynamodb.go
├── types/
│   ├── orders.go
│   ├── fulfillment.go
│   ├── payment.go
```


```
HTTP Client sends JSON:
json
Copy
Edit
{
  "item": "Laptop",
  "quantity": 2
}
gRPC Client sends Protobuf:
proto
Copy
Edit
message CreateOrderRequest {
  string item = 1;
  int32 quantity = 2;
}
```

✅ Handler Layer (OrdersGrpcHandler or OrdersHTTPHandler)
- both call `ordersService.CreateOrder`

✅ Service Layer (OrderService or DatabaseOrderService)

