FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o auth main.go

FROM alpine

WORKDIR /app


COPY --from=builder /app/auth /app/auth

CMD [". /hello"]
