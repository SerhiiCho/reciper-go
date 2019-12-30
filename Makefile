.PHONY: dev
dev:
	cd backend && go run main.go

.PHONY: build
build:
	docker build -t serhiicho/reciper_go:nginx ./frontend
	docker build -t serhiicho/reciper_go:go ./backend
	docker push serhiicho/reciper_go:go
	docker push serhiicho/reciper_go:nginx

.PHONY: test
test:
	go test -cover ./...
	go vet ./...

.DEFAULT_GOAL := dev
