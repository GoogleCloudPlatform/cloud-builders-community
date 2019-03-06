#!/bin/bash

# Checking required inputs
: ${1?"Usage: sync.sh [SOURCE_PATH] [OUTPUT_PATH] [WAITING_SECONDS]"}

[ `gcloud auth list --filter=status:ACTIVE --format="value(account)" 2> /dev/null` ] || { 
  [ -z "${GOOGLE_APPLICATION_CREDENTIALS}" ] && { 
    echo "Error: If no account is activated, GOOGLE_APPLICATION_CREDENTIALS environment variable must be set!" 1>&2
    exit 1
  }
}

[ -f "${GOOGLE_APPLICATION_CREDENTIALS}" ] && {
  gcloud auth activate-service-account --key-file=$GOOGLE_APPLICATION_CREDENTIALS
}

sourcePath=$1
outputPath="${2:-data}" # Default to local data directory, can be overriden
waitingSeconds="${3:-5}" # Default to 5 seconds, can be overriden

# If outputPath does not begin with gs:// create a local directory if doesn not exist
[[ ${outputPath} =~ ^[a-z]+:.*$ ]] || {
  mkdir -p ${outputPath}
}

# Periodically perfom gcs rsync between source and destination paths
while true; do
  echo "Syncing ${sourcePath} every ${waitingSeconds} seconds"
  gsutil -m rsync -d -r ${sourcePath} ${outputPath} || {
    rm -rf /tmp/healthy && break
  }
  echo "Finish synchronizing!"
  touch /tmp/healthy;
  sleep ${waitingSeconds}
done
