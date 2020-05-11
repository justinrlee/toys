#!/bin/sh
while :; do

  echo "$(date): Generating token"
  /app/gcp-exec-creds > /tmp/gcloud/auth_token.tmp
  
  if [[ $? -ne 0 ]]; then
    # If there's an error, then do not copy token file and
    # hope that problem corrects itself before current token expires
    echo "$(date): Error: $(cat /tmp/gcloud/auth_token.tmp)"
  else
    # Should be atomic (no broken file handlers):
    (
      flock -x 200
      mv /tmp/gcloud/auth_token.tmp /tmp/gcloud/auth_token
    ) 200>/tmp/gcloud/auth_token.lock
  fi

  # Cannot find this documented, but have observed that
  # a new token is available 5 minutes before the current one expires
  # so need to re-generate more frequently than that
  sleep 60
done
