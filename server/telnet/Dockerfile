FROM golang:1.12 AS build

ADD . /src

WORKDIR /src

RUN go build -o server

FROM alpine

WORKDIR /app

COPY --from=build /src/server /app/

ENTRYPOINT ./server

