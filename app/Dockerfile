FROM golang:1.14

ENV GO111MODULE=off

RUN go get -u github.com/gorilla/websocket

WORKDIR /go/src/github.com/eifandevs/main

# CMD ["go", "run", "main.go"]
CMD ["tail", "-f", "/dev/null"]
