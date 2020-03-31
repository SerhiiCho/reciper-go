.PHONY: dev
dev:
	cd backend && go run main.go

.PHONY: build
build:
	docker build -t serhiicho/reciper_go:nginx ./frontend
	docker build -t serhiicho/reciper_go:go ./backend
	docker push serhiicho/reciper_go:go
	docker push serhiicho/reciper_go:nginx

.PHONY: migrateup
migrateup:
	migrate -path db/migrations/ -database "mysql://serhii:111111@tcp(localhost:3306)/reciper?charset=utf8" up

.PHONY: migratedown
migratedown:
	migrate -path db/migrations/ -database "mysql://serhii:111111@tcp(localhost:3306)/reciper?charset=utf8" down

.PHONY: migratecreate
migratecreate:
	migrate create -ext sql -dir db/migrations $(arg)

.PHONY: test
test:
	go test -cover ./...
	go vet ./...

.DEFAULT_GOAL := dev
