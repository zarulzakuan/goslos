# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

ENV HTTP_PROXY="http://2993di:Zashasoza24!@razmelpa001.bp.com:80"
ENV HTTPS_PROXY="http://2993di:Zashasoza24!@razmelpa001.bp.com:80"

RUN go mod download

COPY *.go ./

RUN go build -o /test

EXPOSE 8080

CMD [ "/test" ]