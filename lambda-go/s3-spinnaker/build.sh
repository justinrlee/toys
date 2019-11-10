#!/bin/bash

GOOS=linux go build main.go

zip function.zip main

# aws lambda create-function --function-name s3-to-spinnaker --runtime go1.x --zip-file fileb://function.zip --handler main --role arn:aws:iam::XXX:role/lambda-s3-role

aws lambda update-function-code --function-name s3-to-spinnaker --zip-file fileb://function.zip