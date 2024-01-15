DB_URL_EXAMPLE_AUTH=postgres://root:password@localhost:5432/payment_db?sslmode=disable
DB_URL_EXAMPLE_PAYMENT=postgres://root:password@localhost:5432/payment_information_db?sslmode=disable

clearall: clearforup clearnetwork

createall: network postgres createdbpos

network:
	docker network create --driver bridge postgres_network

clearnetwork:
	docker network rm postgres_network

postgres:
	docker run --name postgres_payment_information --network postgres_network -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
	
clearforup:
	docker stop postgres_payment_information && docker rm postgres_payment_information

	
createdbpos:
	docker exec -it postgres_payment_information createdb --username=root --owner=root payment_information_db

initserver:
	go run cmd/api/main.go
	
# dropdb:
#     docker exec -it postgres_payment_information dropdb payment_db

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

protoCard:
	protoc --proto_path=creditService --go_out=creditService --go_opt=paths=source_relative \
	--go-grpc_out=creditService --go-grpc_opt=paths=source_relative \
	creditService/*.proto

sqlcgen:
	sqlc generate



.PHONY: proto postgres createdbpos clearforup network sqlcgen clearnetwork clearall createall initserver protoCard