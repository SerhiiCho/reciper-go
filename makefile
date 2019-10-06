dev-front:
	cd frontend && yarn serve

dev-back:
	go run backend/main.go

build:
	docker build -t serhiicho/reciper:nginx ./frontend
	docker build -t serhiicho/reciper:go ./backend
	docker push serhiicho/reciper:go
	docker push serhiicho/reciper:nginx

test:
	go test -cover ./...