# Dockerfile
FROM golang:1.23.2-alpine

# Instalar dependências básicas: bash, PostgreSQL client, Redis e git
RUN apk add --no-cache bash postgresql-client redis git

# Criar e definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar go.mod e go.sum para instalar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código para o container
COPY . .

# Porta em que a aplicação rodará (caso use)
EXPOSE 8080

# Iniciar Redis como serviço de cache em segundo plano
RUN redis-server --daemonize yes

# Comando para rodar o código Go
CMD ["go", "run", "main.go"]
