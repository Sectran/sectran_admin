#!/bin/bash
openssl genrsa -out ca.key 4096
#创建根证书请求
openssl req -new -key ca.key -out ca.csr -config openssl.cnf
#生成跟证书
openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt -days 365

# #创建用户私钥
openssl genrsa -out user.key 2048
# #创建用户证书签发请求
openssl req -new -key user.key -out user.csr -config user.cnf
# #签发用户证书
openssl x509 -req -in user.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out user.crt -days 1

#查看签发证书信息
# openssl x509 -in ca.crt -text -noout
# openssl x509 -in user.crt -text -noout

