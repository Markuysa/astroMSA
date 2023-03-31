FROM golang:1.19.3-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o main app/cmd/main/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app .

EXPOSE 9002

CMD ["/app/main"]

