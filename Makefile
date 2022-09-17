protoGo:
	protoc --proto_path=backend/api --go_out=backend/api/pb --go_opt=paths=source_relative --go-grpc_out=backend/api/pb --go-grpc_opt=paths=source_relative backend/api/calendar.proto
protoTs:
	protoc -I=./backend/api --js_out=import_style=commonjs:./backend/api/pb --grpc-web_out=import_style=typescript,mode=grpcwebtext:./backend/api/pb backend/api/calendar.proto