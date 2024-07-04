# Etapa de compilação
FROM golang:latest AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos do módulo e baixar as dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar o código fonte do projeto
COPY . .

RUN chmod -R 777 /app/web/templates

# Compilar o aplicativo
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd

# Etapa de execução
FROM alpine:latest

# Instalar bash
RUN apk add --no-cache bash

# Definir o diretório de trabalho
WORKDIR /root/

# Definir variáveis de ambiente
ENV HOST_NAME=localhost
ENV WS_ENDPOINT=/ws
ENV PORT=8080

# Copiar o binário compilado da etapa de compilação
COPY --from=builder /app/main .
# Copiar o diretório de templates
COPY --from=builder /app/web/templates /root/web/templates

# Expor a porta 8080
EXPOSE 8080

# Executar o aplicativo
CMD ["./main"]
