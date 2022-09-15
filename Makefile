lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.49 golangci-lint run -v

all: lint
	go test ./...
	go test ./... -short -race
	go test ./... -run=NONE -bench=. -benchmem
	env GOOS=linux GOARCH=386 go test ./...
