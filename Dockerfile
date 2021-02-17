#Builder stage
FROM golang:1.15 AS BUILDER
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
ENV CGO_ENABLED=0
RUN make build

#RUN_TIME stage
FROM alpine:latest AS RUN_TIME
COPY --from=BUILDER /go/src/app/.env ./
COPY --from=BUILDER /go/src/app/cmd/main ./
CMD ["./main"]