#!/bin/bash

# Authenticate
/builder/kubectl.bash

echo "Running: ko" "$@" >&2
exec "/ko" "$@"
