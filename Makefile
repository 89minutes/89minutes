proto-gen-gateway:
	protoc -I ./proto \
  --go_out ./pb --go_opt paths=source_relative \
  --go-grpc_out ./pb --go-grpc_opt paths=source_relative \
  --swagger_out=:swagger \
  --grpc-gateway_out ./pb --grpc-gateway_opt paths=source_relative \
  proto/*.proto