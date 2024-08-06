FROM golang:1.22 AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/redis/main.go

ENV HOST_REDIS_DOCKER=172.17.0.4

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/
COPY --from=builder /app/cmd/redis/.env /app/

CMD ["./main"]