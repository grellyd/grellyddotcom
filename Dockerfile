FROM golang:1.14

WORKDIR /go/src/github.com/grellyd/grellyddotcom
COPY . .

RUN go build

EXPOSE 8080
