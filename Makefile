.PHONY: deps
deps:
	go install github.com/matryer/moq@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

.PHONY: migrate-up
migrate-up:
	migrate -database "${DB_URL}" -path migration up

.PHONY: migrate-down
migrate-down:
	migrate -database "${DB_URL}" -path migration down 1
