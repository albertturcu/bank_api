buildServer:
	go build -o cmd/server cmd/server/main.go
buildWeb:
	go build -o cmd/web cmd/web/main.go
install:
	go install
dependency:
	go get -d -v ./...
clean:
	go clean
test:
	go test -v -cover ./pkg/...
server:
	nodemon --exec  go run cmd/server/main.go --signal SIGTERM
web:
	nodemon --exec go run cmd/web/main.go --signal SIGTERM