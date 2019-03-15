#!/usr/bin/env bash

set -eou pipefail

#echo "Hello World" | tr 'a-z' 'A-Z'
echo "Hello World" | tr '[:lower:]' '[:upper:]'
