FROM golang:1.19-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./client ./cmd/client/main.go

FROM alpine:3.18 AS app
WORKDIR /app
COPY --from=builder --chown=nobody:nobody ./build/client ./client
RUN chmod +x ./client
# switch to non-root user for security
USER nobody
ENTRYPOINT ["./client"]
