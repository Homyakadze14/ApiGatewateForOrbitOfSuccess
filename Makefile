gen-proto:
	protoc -I proto proto/auth/*.proto --go_out=proto/gen/ --go_opt=paths=source_relative --go-grpc_out=proto/gen/ --go-grpc_opt=paths=source_relative

migrate:
	go run cmd/migrator/main.go -storage-url="postgres://postgres:postgres@localhost:5432/authservice" -migrations-path=./migrations

gen-docs:
	swag init -g=internal/controller/rest/v1/router.go -o=docs --parseInternal