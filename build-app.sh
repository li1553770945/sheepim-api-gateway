#!/bin/bash
RUN_NAME=sheepim-api-gateway
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
go mod tidy
go build -o output/bin/${RUN_NAME}