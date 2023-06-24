# Build stage
FROM golang:1.19.2-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main app/cmd/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
ENV ENV prod
COPY --from=builder /app/main .
COPY --from=builder /app/app.prod.env .

EXPOSE 8080
CMD ["/app/main"]