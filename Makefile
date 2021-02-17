build:
	go build -o cmd/ cmd/main.go
install:
	go install
dependency:
	go get ./...
clean:
	go clean
test:
	go test -v -cover ./pkg/...
run:
	go run cmd/main.go