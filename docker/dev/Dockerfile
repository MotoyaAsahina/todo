FROM golang:1.17.4-alpine as build

WORKDIR /go/src/github.com/MotoyaAsahina/todo

RUN apk add --update --no-cache gcc build-base

ENV DOCKERIZE_VERSION v0.6.1
RUN go install github.com/jwilder/dockerize@$DOCKERIZE_VERSION

RUN go install github.com/cosmtrek/air@v1.27.3

COPY go.* .
RUN go mod download
