FROM golang:1.13-alpine

RUN mkdir -p /app

WORKDIR /app

COPY . /app

RUN go build

CMD ["/app/reciper"]