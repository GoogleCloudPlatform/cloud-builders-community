#!/bin/bash

if [[ $# -ne 2 ]]; then
  echo "Usage: $0 <dir> <tgz>"
  echo "       <dir> is a local directory which will be zip/tarred into a tgz file"
  echo "       <tgz> is a GCS bucket object into which the tgz file will be copied"
  exit 1
fi

cd "$1" || exit
TMP=$(mktemp -d)
localfile="${TMP}/tmp.tgz"
tar -czvf "${localfile}" . || exit
gsutil cp "${localfile}" "$2" || exit

rm -rf ${TMP}
