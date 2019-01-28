#!/usr/bin/env bash

usage()
{
  echo "Usage: $0 <-s|-r> <-f file> [-c output]"
  echo "GCS_BUCKET env variable must be set."
  exit 2
}

set_variable()
{
  local var=$1
  shift
  if [ -z "${!var}" ]; then
    eval "${var}=\"$@\""
  else
    echo "Error: $var already set"
    usage
  fi
}

save_cache()
{
  echo "Saving cache to $GCS_BUCKET/$COMPRESSED"
  zip -qr ${COMPRESSED} ${FILE[*]}
  if [ $? == 0 ]; then
    echo "Uploading to the bucket..."
    gsutil -o GSUtil:parallel_composite_upload_threshold=${THRESHOLD}M cp ${COMPRESSED} ${GCS_BUCKET}/${COMPRESSED}
  else
    echo "Error while zipping!"
  fi
}

restore_cache()
{
  echo "Restoring cache from $GCS_BUCKET/$COMPRESSED"
  gsutil cp ${GCS_BUCKET}/${COMPRESSED} ${COMPRESSED}
  if [ $? == 0 ]; then
    echo "Unzipping..."
    unzip -o -q ${COMPRESSED}
  else
    echo "Couldn't copy $COMPRESSED from the bucket"
    echo "This is perfectly normal for the first usage"
  fi
}

#########################
# Main script starts here

while getopts 'srf:c:t:' opt
do
  case ${opt} in
    s) set_variable ACTION SAVE ;;
    r) set_variable ACTION RESTORE ;;
    f) FILE+=("$OPTARG") ;;
    c) set_variable COMPRESSED $OPTARG ;;
    t) set_variable THRESHOLD $OPTARG ;;
  esac
done

[ -z "$ACTION" ] && usage

[ -z "$FILE" ] && [ "$ACTION" == "SAVE" ] && usage

[ -z "$GCS_BUCKET" ] && usage

if [ -z "$COMPRESSED" ]; then
  COMPRESSED="cache.zip"
fi

if [ -z "THRESHOLD" ]; then
  THRESHOLD=50
fi

case ${ACTION} in
  SAVE) save_cache ;;
  RESTORE) restore_cache ;;
esac

echo "Exiting with 0 - we don't want to break the CloudBuild chain"
exit 0
