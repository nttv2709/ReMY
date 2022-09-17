protoGo:
	protoc --proto_path=backend/api --go_out=backend/api/pb --go_opt=paths=source_relative --go-grpc_out=backend/api/pb --go-grpc_opt=paths=source_relative backend/api/calendar.proto