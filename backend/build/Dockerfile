FROM golang:1.18.3-alpine3.16 AS builder
WORKDIR /app
COPY ./backend .
RUN go build -o main cmd/main.go

FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/main .

COPY wait-for-mysql.sh /wait-for-mysql.sh
RUN chmod +x /wait-for-mysql.sh
RUN apk add --no-cache bash

CMD ["/wait-for-mysql.sh", "mysql-server:3306", "--", "/app/main"]
