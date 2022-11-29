# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY goslos ./

EXPOSE 8080

CMD [ "/gosloss" ]