build:
	go build -o cmd/ cmd/main.go
install:
	go install
clean:
	go clean
test:
	go test -V
run:
	go run cmd/main.go