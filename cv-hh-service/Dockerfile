FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o hh-service

EXPOSE 8082
CMD ["./hh-service"]