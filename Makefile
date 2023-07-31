all: main publisher

main:
	go mod download
	docker compose build
	docker compose up -d
	go run cmd/main.go

publisher:
	go run script/main.go

down:
	docker compose down