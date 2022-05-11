# Build stage
FROM golang:1.18.0-alpine3.14 as builder

WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.14

WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
EXPOSE 8080 8080
CMD ["/app/main"]
