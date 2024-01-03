DB_URL_EXAMPLE=postgres://root:password@localhost:5432/payment_db?sslmode=disable

network:
	docker network create --driver bridge postgres_network

postgres:
	docker run --name postgres_payment --network postgres_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
	
clearforup:
	docker stop postgres_payment && docker rm postgres_payment

	
createdbpos:
	docker exec -it postgres_payment createdb --username=root --owner=root payment_db
	
# dropdb:
#     docker exec -it postgres_payment dropdb payment_db

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

sqlcgen:
	sqlc generate



.PHONY: proto postgres createdbpos clearforup network sqlcgen