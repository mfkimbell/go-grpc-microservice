run-orders:
	@go run ./services/orders

run-kitchen:
	@go run ./services/kitchen

gen:
	@protoc \
		--proto_path=protobuf \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
  		--go-grpc_out=services/common/genproto/orders --go-grpc_opt=paths=source_relative \
		protobuf/orders.proto

.PHONY: run-orders run-kitchen gen
