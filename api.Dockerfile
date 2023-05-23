FROM golang:1.20-alpine AS builder

WORKDIR /cmd

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the api server
CMD ["go", "run", "./cmd/client"]

