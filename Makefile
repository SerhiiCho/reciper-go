.PHONY: build
dev:
	cd backend && go run main.go serve

.DEFAULT_GOAL := dev

build:
	docker build -t serhiicho/reciper:nginx ./frontend
	docker build -t serhiicho/reciper:go ./backend
	docker push serhiicho/reciper:go
	docker push serhiicho/reciper:nginx

test:
	go test -cover ./...