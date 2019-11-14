#!/bin/sh
while :; do

  echo "$(date): Generating token"
  /app/gcp-exec-creds > /tmp/gcloud/auth_token.tmp

  # Should be atomic (no broken file handlers):
  (
    flock -x 200
    mv /tmp/gcloud/auth_token.tmp /tmp/gcloud/auth_token
  ) 200>/tmp/gcloud/auth_token.lock

  sleep 600
done