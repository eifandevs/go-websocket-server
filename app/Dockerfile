FROM golang:1.14

ENV GO111MODULE=off

RUN apt-get update \
 && apt-get -y install default-mysql-client \
 && go get -u github.com/labstack/echo/... \
 && go get -u github.com/labstack/echo/middleware \
 && go get -u github.com/jinzhu/gorm \
 && go get -u github.com/jinzhu/gorm/dialects/mysql \
 && go get -u github.com/BurntSushi/toml \
 && go get -u github.com/abbot/go-http-auth \
 && go get -u github.com/thoas/go-funk



WORKDIR /go/src/github.com/eifandevs/main

# CMD ["go", "run", "main.go"]
CMD ["tail", "-f", "/dev/null"]
