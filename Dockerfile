# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /sl-monitor

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

ARG TRAFFIC_API_AUTH_KEY
ARG MAIL_PASSWORD

COPY --from=build /sl-monitor /sl-monitor
COPY config.yml config.yml
COPY test.db test.db
COPY assets assets

EXPOSE 4444

USER nonroot:nonroot

ENTRYPOINT ["/sl-monitor"]