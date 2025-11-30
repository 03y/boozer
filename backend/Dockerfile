# build
FROM golang:1.23.2-alpine AS builder

WORKDIR /boozer

COPY go.mod ./
COPY go.sum ./

COPY . .

# compile go
#RUN go mod tidy
#RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o /boozer/boozer-backend .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    /bin/sh -c 'go env -w GOPROXY=https://proxy.golang.org,direct && go mod download'

# Copy the rest of the sources
COPY . .

# Build using the same caches so incremental builds reuse state
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /boozer/boozer-backend

# create private key
RUN apk add --no-cache openssl
RUN openssl ecparam -name prime256v1 -genkey -noout -out boozer.pem

# run
FROM alpine:latest

RUN addgroup -S boozergroup && adduser -S boozeruser -G boozergroup

WORKDIR /home/boozeruser

COPY --from=builder /boozer/boozer-backend .
COPY --from=builder /boozer/boozer.pem .

RUN chown boozeruser:boozergroup boozer.pem

# change user
USER boozeruser

EXPOSE 6969

ENTRYPOINT ["./boozer-backend", "0.0.0.0:6969"]

