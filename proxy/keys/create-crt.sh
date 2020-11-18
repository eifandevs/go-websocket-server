#!/bin/sh

# CAの秘密鍵発行
openssl genrsa -out ca-privatekey.pem 2048
# CAの証明書用の署名要求作成
openssl req -new -key ca-privatekey.pem -out ca-csr.pem -sha256 -subj "/C=JP/ST=Tokyo/L=Chuo-ku/O=Local/OU=Local/CN=127.0.0.1"
# CAの証明書発行
openssl req -x509 -key ca-privatekey.pem -in ca-csr.pem -out ca-crt.pem -days 3560
# CAのp12を作成
openssl pkcs12 -export -in ca-crt.pem -inkey ca-privatekey.pem -out ca-crt.p12
# サーバ証明書用の秘密鍵発行
openssl genrsa -out server-privatekey.pem
# サーバ証明書用の署名要求発行
openssl req -new -key server-privatekey.pem -out server-csr.pem -sha256 -subj "/C=JP/ST=Tokyo/L=Chuo-ku/O=Local/OU=Local/CN=127.0.0.1"
# 署名要求にCAの証明書と秘密鍵で署名する
openssl x509 -req -CA ca-crt.pem -CAkey ca-privatekey.pem -CAcreateserial -in server-csr.pem -out server-crt.pem -days 3650 -sha256 -extensions SAN -extfile ./subjectnames.txt

chmod 400 server-crt.pem
