FROM golang:1.24.2-alpine3.21 AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates
EXPOSE 3030
CMD ["./main"]