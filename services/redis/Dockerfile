FROM golang:1.22.2-bullseye AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o uni main.go

FROM alpine
LABEL authors="Stolnikov M."
RUN apk add git make libc6-compat
WORKDIR /app
COPY --from=builder /app/uni /app/uni
COPY broker/config.json /app/redis-connector.json