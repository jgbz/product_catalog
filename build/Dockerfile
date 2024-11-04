FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN env GOOS=linux GOARCH=amd64 go build -o product_catalog_api ./cmd/main.go

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/product_catalog_api .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/static ./static
RUN chmod +x /app/product_catalog_api
EXPOSE 8080
CMD ["./product_catalog_api"]