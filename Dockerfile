# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY /app/main.go .

RUN go build -o ./bin/ .

EXPOSE 8080

RUN echo $(ls ./bin)
ENTRYPOINT [ "./bin/reverse-proxy" ]
