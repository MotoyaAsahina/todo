# syntax = docker/dockerfile:experimental

FROM golang:1.17.4-alpine as server-build
WORKDIR /go/src/github.com/MotoyaAsahina/todo

RUN apk add --update --no-cache gcc build-base

COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download

COPY . .
RUN go build -o /todo


FROM alpine:3.15.0
WORKDIR /app

RUN apk --update --no-cache add tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && mkdir -p /usr/share/zoneinfo/Asia \
  && ln -s /etc/localtime /usr/share/zoneinfo/Asia/Tokyo

COPY --from=server-build /todo ./
