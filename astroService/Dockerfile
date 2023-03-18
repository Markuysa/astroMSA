FROM golang:alpine AS builder

WORKDIR /app

RUN ls

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main app/cmd/main/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 9002
