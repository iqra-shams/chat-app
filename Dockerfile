# Use the official Golang image as a builder stage
FROM golang:latest as builder

LABEL maintainer="Iqra Shams <iqra.shams339@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN go build -o main .





CMD ["./main"]
