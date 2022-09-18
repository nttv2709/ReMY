protoGo:
	protoc --proto_path=backend/api --go_out=backend/api/pb --go_opt=paths=source_relative --go-grpc_out=backend/api/pb --go-grpc_opt=paths=source_relative backend/api/calendar.proto
protoTs:
	protoc -I=./backend/api --js_out=import_style=commonjs:./backend/api/pb --grpc-web_out=import_style=typescript,mode=grpcwebtext:./backend/api/pb backend/api/calendar.proto

deploy:
	docker compose down
	docker compose build
	docker compose up
	# docker tag 4cbc4484f6da registry.heroku.com/dry-oasis-02282/web
push-deploy:
	docker push registry.heroku.com/dry-oasis-02282/web
	heroku container:release web --app=dry-oasis-02282
	heroku logs --tail