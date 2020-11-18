FROM nginx:1.14

ARG node_version="8.x"

# vim
RUN apt-get update && apt-get install -y \
  vim \
  procps \
  iputils-ping \
  net-tools \
  git \
  curl \
  locales \
  openssl

# default locale to ja_JP.UTF-8
RUN locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LC_CTYPE ja_JP.UTF-8
RUN localedef -f UTF-8 -i ja_JP ja_JP.utf8

# timezone
ENV TZ Asia/Tokyo

WORKDIR /var/www

# self-signed certificate
COPY ./keys/server-crt.pem /etc/nginx/server-crt.pem
COPY ./keys/server-privatekey.pem /etc/nginx/server-privatekey.pem

LABEL maintainer "eifandevs <eifan.devs@gmail.com>"