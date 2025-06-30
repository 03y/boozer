# build
FROM golang:1.23.2-alpine AS builder

WORKDIR /boozer

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o /boozer/boozer-backend .

RUN openssl ecparam -name prime256v1 -genkey -noout -out boozer.pem

# run
FROM alpine:latest

RUN addgroup -S boozergroup && adduser -S boozeruser -G boozergroup

WORKDIR /home/boozeruser

COPY --from=builder /boozer/boozer-backend .
COPY --from=builder /boozer/boozer.pem .

# change user
USER boozeruser

EXPOSE 6969

ENTRYPOINT ["./boozer-backend", "0.0.0.0:6969"]

