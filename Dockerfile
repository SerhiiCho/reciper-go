FROM golang:1.13-stretch

RUN mkdir -p /app

WORKDIR /app

COPY . /app

RUN apt-get update \
    && apt-get install gb -y \
    && gb build

CMD ["/app/bin/main"]