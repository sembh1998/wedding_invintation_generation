# Build stage
FROM golang:1.21-alpine3.17 as  builder
WORKDIR /app
COPY . .
RUN go build -o main .\cmd\httpserver\main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 80
CMD ["/app/main"]
