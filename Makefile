BINARY_NAME=customer_service

build:
	go build -o ${BINARY_NAME}-windows ./cmd/main.go

run: build
	${BINARY_NAME}-windows

testing_db:
	docker-compose up --build -d

test: testing_db
	go test ./...
