#!/bin/bash

# export path to the locally installed firebase tools
export PATH=$PATH:/node_modules/firebase-tools/bin

# run the original firebase
firebase "$@" --token $FIREBASE_TOKEN