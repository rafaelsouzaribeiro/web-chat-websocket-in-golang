Web chat with WebSocket, Redis and Cassandra, including notifications for logged-in and logged-out users, and emoji support, implemented in Go and JavaScript.
<br /><br />
To use messages and track connected and disconnected users with only a map variable, use this project.<a href="https://github.com/rafaelsouzaribeiro/Web-chat-with-WebSocket-using-a-map-variable-in-Go">here</a>.<br /><br />

If you want to use a chat with message tracking and the ability to track connected and disconnected users using Redis or Cassandra, this is the project you need.<br />
<br />
1 - Navigate to the /cmd directory.<br/>
2 - Set within the cmd/main.go code: `f := factory.NewFactory(factory.Cassandra, Conf)`, using either factory.Redis or factory.Cassandra.<br/>
3 - Run: main.go<br />
4 - access via browser: http://localhost:8080/chat<br />
5 - being able to open in multiple tabs and connect multiple users<br />
6 - connect user and send message

<br/>

You can also run it through the dockerfile. I am using Docker Engine version 27.1.2:
<br/>
Redis:<br/>

 ```
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/main.go

FROM scratch
WORKDIR /app

ENV HOST_REDIS_DOCKER="172.17.0.4"

COPY --from=builder /app/main /app/
COPY --from=builder /app/cmd/.env /app/

CMD ["./main"]

 ```
 <br />
 Cassandra:
  <br />

 ```
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/main.go

FROM scratch
WORKDIR /app

ENV HOST_CASSANDRA_DOCKER="172.17.0.4"

COPY --from=builder /app/main /app/
COPY --from=builder /app/cmd/.env /app/

CMD ["./main"]

 ```
<br />
To create the keyspace and tables in Cassandra, just run the following Makefile commands in the project root:
<br/><br/>

 ```
Create KeysPace: make create-keyspace
Create tables: make migrateup
Drop tables and keyspace: make migratedown

 ```
If you don't have make installed, install it with the following command:
```
sudo apt install make

 ```

To run Redis in Docker, navigate to the internal/infra/database/redis directory and run:<br/>
 ```
 sudo docker compose up
 ```
<br />
If you want to use Cassandra, navigate to the internal/infra/database/cassandra directory and execute the same command.<br/><br/>
 
 To use Redis and Cassandra as Docker containers and access them from another WebSocket container, you need to determine the internal IP address of the Redis or Cassandra container. First, with Redis or Cassandra running, execute the following command with the container name:<br/>

 ```
sudo docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis
 ```


 ```
sudo docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' cassandra
 ```
<br/>
 
 After running this command, you can set the Redis or Cassandra IP address in the Dockerfile using the ENV instruction:

```
ENV HOST_CASSANDRA_DOCKER=ip_address_from_inspect

 ```

```
ENV HOST_REDIS_DOCKER=ip_address_from_inspect

 ```






