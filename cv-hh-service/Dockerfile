FROM golang:1.20 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8082
CMD ["./app"]