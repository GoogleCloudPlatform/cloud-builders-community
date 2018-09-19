#!/bin/bash

# If there is no current context, get one.
if [[ $(kubectl config current-context 2> /dev/null) == "" ]]; then
    # This tries to read environment variables. If not set, it grabs from gcloud
    cluster=${CLOUDSDK_CONTAINER_CLUSTER:-$(gcloud config get-value container/cluster 2> /dev/null)}
    region=${CLOUDSDK_COMPUTE_REGION:-$(gcloud config get-value compute/region 2> /dev/null)}
    zone=${CLOUDSDK_COMPUTE_ZONE:-$(gcloud config get-value compute/zone 2> /dev/null)}
    project=${GCLOUD_PROJECT:-$(gcloud config get-value core/project 2> /dev/null)}

    function var_usage() {
        cat <<EOF
No cluster is set. To set the cluster (and the region/zone where it is found), set the environment variables
  CLOUDSDK_COMPUTE_REGION=<cluster region> (regional clusters)
  CLOUDSDK_COMPUTE_ZONE=<cluster zone> (zonal clusters)
  CLOUDSDK_CONTAINER_CLUSTER=<cluster name>
EOF
        exit 1
    }

    [[ -z "$cluster" ]] && var_usage
    [ ! "$zone" -o "$region" ] && var_usage

    if [ -n "$region" ]; then
      echo "Running: gcloud config set container/use_v1_api_client false"
      gcloud config set container/use_v1_api_client false
      echo "Running: gcloud beta container clusters get-credentials --project=\"$project\" --region=\"$region\" \"$cluster\""
      gcloud beta container clusters get-credentials --project="$project" --region="$region" "$cluster" || exit
    else
      echo "Running: gcloud container clusters get-credentials --project=\"$project\" --zone=\"$zone\" \"$cluster\""
      gcloud container clusters get-credentials --project="$project" --zone="$zone" "$cluster" || exit
    fi
fi

echo "Running: helm init --client-only"
helm init --client-only

# check if repo values provided then add that repo
if [[ -n $HELM_REPO_NAME && -n $HELM_REPO_URL ]]; then
  echo "Adding chart helm repo $HELM_REPO_URL "
  helm repo add $HELM_REPO_NAME $HELM_REPO_URL
fi

echo "Running: helm repo update"
helm repo update

# check if Tillerless value 'TILLERLESS=true' is provided then install and start the plugin
if [ "$TILLERLESS" = true ]; then
  echo "Installing Tillerless plugin"
  helm plugin install https://github.com/rimusz/helm-tiller
  echo "Starting Tillerless plugin"
  helm tiller start-ci "$TILLER_NAMESPACE"
  echo
  export HELM_HOST=localhost:44134
  if [ "$DEBUG" = true ]; then
      echo "Running: helm $@"
  fi
  helm "$@"
  helm tiller stop
else
  if [ "$DEBUG" = true ]; then
      echo "Running: helm $@"
  fi
  helm "$@"
fi
