FROM golang:1.25

WORKDIR /go/src/github.com/grellyd/grellyddotcom
COPY . .

RUN go build

EXPOSE 8080
EXPOSE 8443
