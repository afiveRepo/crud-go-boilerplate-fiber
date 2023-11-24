FROM golang:latest

WORKDIR /app

RUN go mod tidy

EXPOSE 8080

CMD go run main.go