FROM golang:1.13-alpine

COPY . /app

RUN cd /app && go build -o reciper

CMD ["/app/reciper"]