#! /bin/bash

go build -o tdg ./cmd/tdg/main.go && ./tdg ./sample.json 100 > result.txt
