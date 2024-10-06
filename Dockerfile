FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM builder AS production
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tg-game .

FROM alpine:latest

COPY --from=production /app/myapp /usr/local/bin/myapp

CMD ["tg-game"]

ENV ENV=production

EXPOSE 8080
