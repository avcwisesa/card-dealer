run:
	go run main.go

test: format
	@./scripts/coverage.sh
	go tool cover -html=coverage.out -o coverage.html

format:
	go fmt ./...
