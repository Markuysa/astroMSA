FROM golang:alpine AS builder

WORKDIR /app

COPY auth .

RUN go build -o auth main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/auth /app/auth

