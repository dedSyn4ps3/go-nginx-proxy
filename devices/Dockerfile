FROM golang:1.17.2-buster

WORKDIR /go/src/app
COPY . .

RUN go mod init awesome.inc/device-api
RUN go mod tidy

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/sirupsen/logrus

RUN apt update
RUN apt install -y python3

CMD ["go", "run", "device-api.go"]