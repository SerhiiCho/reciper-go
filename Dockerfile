FROM golang:1.13-alpine

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

CMD ["go-wrapper", "run"]

ONBUILD COPY ./src /go/src/app
ONBUILD RUN go-wrapper download
ONBUILD RUN go-wrapper install