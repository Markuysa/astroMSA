FROM golang:1.19.3-alpine AS builder

WORKDIR /appdir

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o main app/cmd/main/main.go

FROM alpine

WORKDIR /appdir
# Здесь попробовать уменьшить размер образа через COPY --from=builder /appdir main
# Здесь попробовать уменьшить размер образа через COPY --from=builder /appdir configdir
COPY --from=builder /appdir .


CMD ["/appdir/main"]


