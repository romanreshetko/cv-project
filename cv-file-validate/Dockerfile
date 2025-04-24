FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o validate-app

EXPOSE 8081
CMD ["./validate-app"]