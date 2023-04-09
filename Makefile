
protogateway:
	protoc --proto_path=messageSenderService/app/protobuf --go_out=apiGateway/app/services/pb --go_opt=paths=source_relative \
        --go-grpc_out=apiGateway/app/services/pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=apiGateway/app/services/pb --grpc-gateway_opt=paths=source_relative \
         messageSenderService/app/protobuf/*.proto
	protoc --proto_path=authService/app/protobuf --go_out=apiGateway/app/services/pb --go_opt=paths=source_relative \
            --go-grpc_out=apiGateway/app/services/pb --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=apiGateway/app/services/pb --grpc-gateway_opt=paths=source_relative \
             authService/app/protobuf/*.proto
	protoc --proto_path=astroService/app/protobuf --go_out=apiGateway/app/services/pb --go_opt=paths=source_relative \
            --go-grpc_out=apiGateway/app/services/pb --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=apiGateway/app/services/pb --grpc-gateway_opt=paths=source_relative \
             astroService/app/protobuf/*.proto
protoAstro:
	protoc --proto_path=astroService/app/protobuf --go_out=astroService/app/protobuf/pb --go_opt=paths=source_relative \
                --go-grpc_out=astroService/app/protobuf/pb --go-grpc_opt=paths=source_relative \
                --grpc-gateway_out=astroService/app/protobuf/pb --grpc-gateway_opt=paths=source_relative \
                 astroService/app/protobuf/*.proto
protoMessages:
	protoc --proto_path=messageSenderService/app/protobuf --go_out=messageSenderService/app/protobuf/pb --go_opt=paths=source_relative \
			   --go-grpc_out=messageSenderService/app/protobuf/pb --go-grpc_opt=paths=source_relative \
			   --grpc-gateway_out=messageSenderService/app/protobuf/pb --grpc-gateway_opt=paths=source_relative \
				messageSenderService/app/protobuf/*.proto
protoAuth:
	protoc --proto_path=authService/app/protobuf --go_out=authService/app/protobuf/pb --go_opt=paths=source_relative \
			   --go-grpc_out=authService/app/protobuf/pb --go-grpc_opt=paths=source_relative \
			   --grpc-gateway_out=authService/app/protobuf/pb --grpc-gateway_opt=paths=source_relative \
				authService/app/protobuf/*.proto
some:
	protoc --proto_path=authService/app/protobuf/google/type --go_out=authService/app/protobuf/pb --go_opt=paths=source_relative \
           --go-grpc_out=authService/app/protobuf/pb --go-grpc_opt=paths=source_relative authService/app/protobuf/google/type/date.proto
build:
	go build -o astroService/main.exe astroService/app/cmd/main/main.go
	go build -o authService/main.exe authService/app/cmd/main/main.go
	go build -o messageSenderService/main.exe messageSenderService/app/cmd/main/main.go


astroProto:
	protoc --proto_path=apiGateway/app/protobuf --go_out=apiGateway/app/protobuf/gen --go_opt=paths=source_relative \
			--go-grpc_out=apiGateway/app/protobuf/gen --go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=apiGateway/app/protobuf/gen --grpc-gateway_opt=paths=source_relative \
			 apiGateway/app/protobuf/astroServices.proto

PHONY proto protoAstro protoMessages protoAuth:

