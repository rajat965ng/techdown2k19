#!/usr/bin/env bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.