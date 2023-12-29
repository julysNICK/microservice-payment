DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto


.PHONY: proto