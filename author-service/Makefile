.PHONY: all

SHELL := /bin/bash # Use bash syntax
SOURCE_DIRS=$(shell go list ./... | grep -v /vendor | grep -v /out | cut -d "/" -f2 | sort | uniq)

IN_EXE="./cmd"
OUT_EXE="./bin/report-service"

POSTGRESQL_URL="postgresql://user:pass@localhost:5435/authordb?sslmode=disable"

build:
	@GO111MODULE=on go build -o $(OUT_EXE) $(IN_EXE)

start-server: build
	@$(OUT_EXE)

test:
	@GO111MODULE=on go test ./...

proto:
	protoc --go_out=./authorpb --go-grpc_out=./authorpb ./authorpb/author.proto

migrate-up:
	docker run -v /home/projects/Test/synapsis/author-service/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database ${POSTGRESQL_URL} up

migrate-down:
	docker run -v /home/projects/Test/synapsis/author-service/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database ${POSTGRESQL_URL} down 1

create-migration:
    @read -p "Enter migration name: " name; \
    migrate create -ext sql -dir ./migrations -seq $$name
