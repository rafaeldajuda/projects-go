# stage 1 - build
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api

# stage 2 - runtime
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3000
CMD ["./app"]