run:
	go run main.go

test: format
	go test ./...

format:
	go fmt ./...