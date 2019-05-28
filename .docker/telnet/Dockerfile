FROM golang:1.12 as build

ADD . /src

WORKDIR /src/server/telnet

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o telnet-server

#ENTRYPOINT ./telnet-server -addr=0.0.0.0:4000

FROM alpine

WORKDIR /app

COPY --from=build /src/server/telnet/telnet-server /app/

EXPOSE 4000

ENTRYPOINT ./telnet-server -addr=0.0.0.0:4000

