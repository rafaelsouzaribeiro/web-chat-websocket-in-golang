Web chat with WebSocket and Redis, including notifications for logged-in and logged-out users, and emoji support, implemented in Go and JavaScript."
<br />
<br />
1 - Run: cmd/main.go<br />
2 - access via browser: http://localhost:8080/chat<br />
2 - being able to open in multiple tabs and connect multiple users<br />
3 - connect user and send message

<br/>

You can also run it through the dockerfile:<br />

 ```
FROM golang:1.22.0 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

ENV HOST_NAME=0.0.0.0
ENV WS_ENDPOINT=/ws
ENV PORT=8080
ENV HOST_REDIS=<ip_address_from_inspect>
ENV PORT_REDIS=6379

COPY --from=builder /app/main .

EXPOSE $PORT

CMD ["./main"]

 ```
 <br />
To use messages and track connected and disconnected users only with a map variable, use version: v1.0.0.<br /><br />

Now, if you want to use chat with messages and track connected and disconnected users with Redis, use version: v1.1.0.<br /><br />
To run Redis in Docker, use:
 ```
sudo docker run --name redis -d -p 6379:6379 redis
 ```
<br />
"To use Redis as a Docker container and access it from another WebSocket container, you need to determine the internal IP address of the Redis container. First, with Redis running, execute the following command: 
 ```
sudo docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis
 ```
 <br/>
 After running this command, you can set the Redis IP address in the Dockerfile using the ENV instruction:<br/>
```
 ENV HOST_REDIS=<ip_address_from_inspect>
 ```



