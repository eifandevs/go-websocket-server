FROM golang:1.14

ENV GO111MODULE=off

RUN go get -u github.com/labstack/echo/... \
 && go get -u github.com/labstack/echo/middleware \
 && go get -u github.com/abbot/go-http-auth \
 && go get -u github.com/thoas/go-funk \
 && go get -u github.com/gorilla/websocket

WORKDIR /go/src/github.com/eifandevs/main

# CMD ["go", "run", "main.go"]
CMD ["tail", "-f", "/dev/null"]
