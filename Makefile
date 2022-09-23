.PHONY: deps
deps:
	go install github.com/matryer/moq@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: migrate-up
migrate-up:
	migrate -database "${DB_URL}" -path migration up

.PHONY: migrate-down
migrate-down:
	migrate -database "${DB_URL}" -path migration down 1
