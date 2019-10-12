FROM golang:1.13-alpine

COPY . /app

WORKDIR /app

RUN cd /app && go build -o reciper

CMD ["./reciper"]