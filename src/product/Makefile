DB_URL=postgresql://minipulsa:minipulsa@localhost:5432/product?sslmode=disable

STATE?=up
migrate:
	migrate -path schema -database "$(DB_URL)" -verbose $(STATE)

test:
	go test -v -cover ./...

run-dev:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

proto:
	protoc --proto_path=../../pb --go_out=./adapter/grpc --go_opt=paths=source_relative \
	--go-grpc_out=./adapter/grpc --go-grpc_opt=paths=source_relative \
	../../pb/*.proto


.PHONY: migrate test run-dev mock proto