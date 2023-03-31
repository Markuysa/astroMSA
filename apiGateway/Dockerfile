FROM golang:1.19.3-alpine AS builder

WORKDIR /app


COPY . ./

RUN go build -o main app/cmd/main/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8085

CMD ["/app/main"]

