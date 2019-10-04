FROM golang:1.13-alpine

RUN mkdir -p /app/app

WORKDIR /app/app

COPY . /app

RUN apk add  --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/v3.7/main/ nodejs=8.9.3-r1 \
    && npm i \
    && npm run build \
    && cd /app && go build

CMD ["/app/reciper"]