#!/bin/bash

if [[ $# -ne 2 ]]; then
  echo "Usage: $0 <tgz> <target>"
  echo "       <tgz> is a tgz object in a GCS bucket"
  echo "       <target> is a local directory into which you want tgz expanded"
  exit 1
fi

TMP=$(mktemp -d)
localfile="${TMP}/tmp.tgz"
gsutil cp "$1" "${localfile}" || exit
mkdir -p "$2"
cd "$2" || exit
tar -xzvf "${localfile}" . || exit

rm -rf "${TMP}"
