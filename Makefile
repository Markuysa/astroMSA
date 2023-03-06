
proto:
	protoc --proto_path=apiGateway/app/services/protobuf --go_out=apiGateway/app/services/pb --go_opt=paths=source_relative \
        --go-grpc_out=apiGateway/app/services/pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=apiGateway/app/services/pb --grpc-gateway_opt=paths=source_relative \
         apiGateway/app/services/protobuf/*.proto
PHONY proto: