FROM golang:alpine AS builder
#прикол в том, что я просто сбилдил бинарь, его нужно еще запустить!
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main app/cmd/main/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 9000

