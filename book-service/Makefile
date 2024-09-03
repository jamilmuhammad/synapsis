.PHONY: all

SHELL := /bin/bash # Use bash syntax
SOURCE_DIRS=$(shell go list ./... | grep -v /vendor | grep -v /out | cut -d "/" -f2 | sort | uniq)

IN_EXE="./cmd"
OUT_EXE="./bin/report-service"

POSTGRESQL_URL="postgresql://user:pass@localhost:5436/bookdb?sslmode=disable"

build:
	@GO111MODULE=on go build -o $(OUT_EXE) $(IN_EXE)

start-server: build
	@$(OUT_EXE)

test:
	@GO111MODULE=on go test ./...

proto:
	protoc --go_out=./bookpb --go-grpc_out=./bookpb ./bookpb/book.proto
	protoc --go_out=./bookpb --go-grpc_out=./bookpb ./bookpb/loan.proto

migrate-up:
	docker run -v /home/projects/Test/synapsis/book-service/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database ${POSTGRESQL_URL} up

migrate-down:
	docker run -v /home/projects/Test/synapsis/book-service/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database ${POSTGRESQL_URL} down 1

create-migration:
    @read -p "Enter migration name: " name; \
    migrate create -ext sql -dir ./migrations -seq $$name
