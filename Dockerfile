FROM golang:1.16-alpine

WORKDIR /app
COPY . .

RUN go get -u
RUN go build .

ENV PASTE_PORT 4000
ENV PASTE_BIND 0.0.0.0
