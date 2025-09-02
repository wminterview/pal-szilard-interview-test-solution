FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum main.go ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o library ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/library .

RUN adduser -D -g '' appuser
USER appuser

CMD ["./library"]