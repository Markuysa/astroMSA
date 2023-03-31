FROM golang:1.19.3-alpine AS builder

WORKDIR /appdir

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o main app/cmd/main/main.go

FROM alpine

WORKDIR /appdir

COPY --from=builder /app .

EXPOSE 9000

CMD ["/app/main"]

