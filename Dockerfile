FROM golang:1.23.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-shortener ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/url-shortener .
COPY --from=builder /app/internal/templates ./internal/templates

EXPOSE 8080

CMD ["./url-shortener"]
