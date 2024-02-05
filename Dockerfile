FROM golang:1.21-alpine as builder
LABEL authors="Maksim"

ENV GOOS linux

WORKDIR /app

COPY ./ .

RUN go mod download

RUN go build -o GroupAssist ./cmd/app/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/GroupAssist /app/GroupAssist
COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/database /app/database

EXPOSE 8080

ENTRYPOINT ["./GroupAssist"]