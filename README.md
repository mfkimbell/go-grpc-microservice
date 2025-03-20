# go-grpc-microservice

#### Architechture
<img width="701" alt="grpc_arch" src="https://github.com/user-attachments/assets/0b15cf03-5cb5-467f-9ff6-09d66f638d94" />

#### Why do we care about gRPC?

1. **Performance (faster in most cases) w/ protocol buffers and http2**
   
Protobuf is a binary format, which is typically smaller and faster to parse than JSON.
gRPC streams data over HTTP/2 by default, allowing faster, multiplexed connections – no need to open multiple TCP sockets.
  
2. **Strongly-Typed Contracts causes less errors**
   
Protobuf (.proto files) define strict schemas for requests and responses.
Each language has auto-generated code from .proto, ensuring type safety and fewer runtime errors.
With HTTP/JSON, you must manually maintain schemas (e.g., OpenAPI), but it’s still just dynamic JSON in the end – more prone to mismatches.

  
3. **Built-In Streaming with HTTP2**
   
gRPC supports client-streaming, server-streaming, and bidirectional streaming out of the box.
HTTP/JSON typically uses long polling or WebSockets for streaming, which can be more complex.
  
4. **Generate matching code for all lanagues based on proto file**
   
gRPC auto-generates client stubs and server interfaces for each language (Go, Python, Java, JavaScript, etc.).
This means less boilerplate and consistent data structures across your entire stack.
HTTP/JSON typically requires manually writing client libraries or using frameworks like OpenAPI, but not as seamlessly integrated.

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



