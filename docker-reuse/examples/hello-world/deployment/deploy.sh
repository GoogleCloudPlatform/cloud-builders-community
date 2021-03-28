#!/bin/sh

echo "Applying hello-world deployment configuration:"

cat "`dirname "$0"`/deployment.yaml"
