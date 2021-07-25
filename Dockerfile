FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 4000

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

RUN go build -o gin-api