dev:
	go run main.go

build:
	docker-compose build
	docker-compose push serhiicho/reciper:go
	docker-compose push serhiicho/reciper:nginx

prod:
	docker-compose up -d

test:
	go test -cover ./...