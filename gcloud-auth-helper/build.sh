#!/bin/bash

DATE=$(date +%s)

docker build -t justinrlee/gcloud-auth-helper:${DATE} .

docker push justinrlee/gcloud-auth-helper:${DATE}