# Build stage
FROM golang:1.19.2-alpine3.16 AS builder
WORKDIR /apps
COPY . .
RUN go build -o main src/main.go

# Run stage
FROM alpine:3.16
WORKDIR /apps
COPY --from=builder /apps/main .
EXPOSE 8080
CMD ["./main"]