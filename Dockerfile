FROM golang:1.13-stretch

RUN mkdir -p /app

WORKDIR /app

COPY . /app

RUN go build

CMD ["/app/reciper"]