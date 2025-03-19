# go-grpc-microservice

#### Setup
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

✅ Handler Layer (OrdersGrpcHandler or OrdersHTTPHandler)
- The handler layer processes the request and calls the service.
- both call `ordersService.CreateOrder`

✅ Service Layer (OrderService)
- interacts with the database
- The service layer does NOT know anything about gRPC or HTTP.

```
services/orders/
├── handler/        # HTTP & gRPC request handling
│   ├── http.go
│   ├── grpc.go
├── service/        # Business logic
│   ├── orders.go
│   ├── payment.go  
├── types/         
│   ├── types.go
```



