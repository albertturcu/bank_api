build:
	go build -o cmd/api cmd/api/main.go
install:
	go install
dependency:
	go get -d -v ./...
clean:
	go clean
test:
	go test -v -cover ./pkg/...
api:
	go run cmd/api/main.go
web:
	go run cmd/web/main.go