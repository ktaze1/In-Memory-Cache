FROM golang:1.17.5

RUN mkdir /build

ADD . /build

WORKDIR /build

RUN go build -o main

CMD ["/build/main"]