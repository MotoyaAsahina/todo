# syntax = docker/dockerfile:experimental
FROM golang:1.17.4-alpine as build

WORKDIR /go/src/github.com/MotoyaAsahina/todo

RUN apk add --update --no-cache gcc build-base

ENV DOCKERIZE_VERSION v0.6.1
RUN --mount=type=cache,target=/root/.cache/go-build \
  go install github.com/jwilder/dockerize@$DOCKERIZE_VERSION

RUN --mount=type=cache,target=/root/.cache/go-build \
  go install github.com/cosmtrek/air@v1.27.3

COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download
