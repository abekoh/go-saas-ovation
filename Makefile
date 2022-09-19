.PHONY: deps
deps:
		go install github.com/matryer/moq@latest
		go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
