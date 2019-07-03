#!/bin/bash

# Colours
RED='\033[0;31m'
RESET='\033[0m'

# Accepted Parameters
HELP_PARAMETER="--help"
NO_BUILD_PARAMETER="--no-build"
YARN_BUILD_SCRIPT_PARAMETER="--yarn-build-script"
NO_TEST_PARAMETER="--no-test"
YARN_TEST_SCRIPT_PARAMETER="--yarn-test-script"
DEPLOYMENT_NAME_PARAMETER="--deployment-name"
RUNTIME_PARAMETER="--runtime"
ENTRY_POINT_PARAMETER="--entry-point"
ENV_VARS_FILE_PARAMETER="--env-vars-file"

# Describe Usage
function usage() {
  if [ -n "$1" ]; then
    echo -e "$RED$1$RESET\n";
  fi
  echo -e "Usage: $0"
  echo -e ""
  echo -e "  --help"
  echo -e "     List parameter descriptions."
  echo -e ""
  echo -e "  --no-build"
  echo -e "     Do not execute a build."
  echo -e ""
  echo -e "  --yarn-build-script=YARN_BUILD_SCRIPT"
  echo -e "     The script defined in the package.json which builds the project."
  echo -e "     Default value: 'build'"
  echo -e ""
  echo -e "  --no-test"
  echo -e "     Do not execute tests."
  echo -e ""
  echo -e "  --yarn-test-script=YARN_TEST_SCRIPT"
  echo -e "     The script defined in the package.json which tests the project."
  echo -e "     Default value: 'test'"
  echo -e ""
  echo -e "  --deployment-name=DEPLOYMENT_NAME ($RED*$RESET)"
  echo -e "     The name that should be used when deploying the Cloud Function."
  echo -e ""
  echo -e "  --runtime=RUNTIME"
  echo -e "     Runtime in which to run the function."
  echo -e "     Options: 'nodejs6' (deprecated), 'nodejs8' or 'nodejs10'"
  echo -e "     Default value: 'nodejs10'"
  echo -e ""
  echo -e "  --entry-point=ENTRY_POINT"
  echo -e "     The named export which should be deployed."
  echo -e "     Default value: The value assigned to '--deployment-name'"
  echo -e ""
  echo -e "  --env-vars-file=ENV_VARS_FILE"
  echo -e "     The file containing environment variables to be supplied to the cloud function."
  echo -e ""
  echo -e "Example: $0 --deployment-name=my-function"
  if [ -n "$1" ]; then
    exit 1
  else
    exit 0
  fi
}

# Set Default Parameters
YARN_BUILD_SCRIPT="build"
YARN_TEST_SCRIPT="test"
RUNTIME="nodejs10"

# Parse Parameters
while [[ "$#" > 0 ]]; do case $1 in
  $HELP_PARAMETER) usage;;
  $NO_BUILD_PARAMETER) NO_BUILD="true"; shift;;
  $YARN_BUILD_SCRIPT_PARAMETER=*) YARN_BUILD_SCRIPT="${1#"$YARN_BUILD_SCRIPT_PARAMETER="}"; shift;;
  $NO_TEST_PARAMETER) NO_TEST="true"; shift;;
  $YARN_TEST_SCRIPT_PARAMETER=*) YARN_TEST_SCRIPT="${1#"$YARN_TEST_SCRIPT_PARAMETER="}"; shift;;
  $DEPLOYMENT_NAME_PARAMETER=*) DEPLOYMENT_NAME="${1#"$DEPLOYMENT_NAME_PARAMETER="}"; shift;;
  $RUNTIME_PARAMETER=*) RUNTIME="${1#"$RUNTIME_PARAMETER="}"; shift;;
  $ENTRY_POINT_PARAMETER=*) ENTRY_POINT="${1#"$ENTRY_POINT_PARAMETER="}"; shift;;
  $ENV_VARS_FILE_PARAMETER=*) ENV_VARS_FILE="${1#"$ENV_VARS_FILE_PARAMETER="}"; shift;;
  *) usage "The '$1' parameter is invalid."; shift;;
esac; done

# Validate Parameters
if [ -n "$DEPLOYMENT_NAME" ]; then usage "The '--deployment-name' parameter is not set."; fi;
if [ "$RUNTIME" != "nodejs6" ] && [ "$RUNTIME" != "nodejs8" ] && [ "$RUNTIME" != "nodejs10" ]; then
  usage "The '--runtime' parameter is invalid.";
fi;

# Install Dependencies
yarn

# Build
if [ -n "$NO_BUILD" ]; then yarn $YARN_BUILD_SCRIPT; fi;

# Test
if [ -n "$NO_TEST" ]; then yarn $YARN_TEST_SCRIPT; fi;

# Deploy
gcloud functions deploy $DEPLOYMENT_NAME \
  --trigger-http \
  --runtime=$RUNTIME \
  ${ENTRY_POINT:+"--entry-point=$ENTRY_POINT"} \
  ${ENV_VARS_FILE:+"--env-vars-file=$ENV_VARS_FILE"}
