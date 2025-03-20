# go-grpc-microservice

#### Architechture
<img width="601" alt="grpc_arch" src="https://github.com/user-attachments/assets/0b15cf03-5cb5-467f-9ff6-09d66f638d94" />


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

Then run:
`make gen` to build the protocol buffer server interfaces code
`make orders` to build the 8000 http server and 9000 gRPC server
`make kitchen` to build the 1000 http frontend server

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



