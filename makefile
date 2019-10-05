dev:
	go run backend/main.go

build:
	docker-compose build
	docker push serhiicho/reciper:go
	docker push serhiicho/reciper:nginx

prod:
	docker-compose up -d

test:
	go test -cover ./...