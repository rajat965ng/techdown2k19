#!/usr/bin/env bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.


# Generate stuff for SSL

#1. Generate certificate authority + trust certificate (ca.crt)
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj "/CN=localhost"   #Shareable


#2. Generate server private key
openssl genrsa -passout pass:1111 -des3 -out server.key 4096


#3. Generate certificate signing request from the CA (server.csr)
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=localhost"  #Shareable

#4. Sign certificate with CA we created (Self signing) - server.crt -(keep on server)
openssl x509 -req -passin pass:1111 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt


#5. Convert server certificate to (.pem) format (server.pem) - used by gRPC
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem