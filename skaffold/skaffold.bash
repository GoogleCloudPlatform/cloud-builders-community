#!/bin/bash
set -e

# Invoke kubectl.bash to process CLOUDSDK_* environment variables and make a
# call to the cluster to verify connectivity.
/builder/kubectl.bash version

echo "Running: skaffold $@"
skaffold "$@"
