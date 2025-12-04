FROM golang:1.24.0-alpine AS builder
WORKDIR /app

ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build only the main entry file
RUN go build -o server ./cmd

FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
