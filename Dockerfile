FROM golang:1.20-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main /app/cmd/server/main.go


FROM alpine:3.18 as dev

WORKDIR /app

COPY .env /app/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]