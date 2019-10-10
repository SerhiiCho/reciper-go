.PHONY: dev
dev:
	cd backend && go run main.go

.PHONY: build
build:
	docker build -t serhiicho/reciper:nginx ./frontend
	docker build -t serhiicho/reciper:go ./backend
	docker push serhiicho/reciper:go
	docker push serhiicho/reciper:nginx

.PHONY: test
test:
	go test -cover -race ./...

.DEFAULT_GOAL := dev
