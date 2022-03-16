# syntax = docker/dockerfile:experimental

FROM golang:1.17.4-alpine as server-build
WORKDIR /go/src/github.com/MotoyaAsahina/todo

RUN apk add --update --no-cache gcc build-base

COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download

COPY . .
RUN go build -o /todo


FROM node:16.13.1-alpine3.14 as client-build
WORKDIR /client

RUN --mount=type=cache,target=/var/cache/apk \
  apk add --update --no-cache python3 make g++

COPY client/package.json client/package-lock.json ./
RUN --mount=type=cache,target=/root/.npm \
  npm ci

COPY client .
RUN --mount=type=cache,target=/client/node_modules/.cache \
  npm run build


FROM alpine:3.15.0
WORKDIR /app

RUN apk --update --no-cache add tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && mkdir -p /usr/share/zoneinfo/Asia \
  && ln -s /etc/localtime /usr/share/zoneinfo/Asia/Tokyo

COPY --from=server-build /todo ./
COPY --from=client-build /client/dist/ ./client/dist/
