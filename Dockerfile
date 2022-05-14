# Build stage
FROM golang:1.18.0-alpine3.14 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/server/main.go

# Run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8000 8000
CMD ["/app/main"]
