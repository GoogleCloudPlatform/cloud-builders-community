# [Helm](https://docs.helm.sh/) tool builder

## Using this builder with Google Container Engine

To use this builder, your
[builder service account](https://cloud.google.com/cloud-build/docs/how-to/service-account-permissions)
will need IAM permissions sufficient for the operations you want to perform. For
typical read-only usage, the "Kubernetes Engine Viewer" role is sufficient. To
deploy container images on a GKE cluster, the "Kubernetes Engine Developer" role
is sufficient. Check the
[GKE IAM page](https://cloud.google.com/container-engine/docs/iam-integration)
for details.

For most use, kubectl will need to be configured to point to a specific GKE
cluster. You can configure the cluster by setting environment variables.

    # Set region for regional GKE clusters or Zone for Zonal clusters
    CLOUDSDK_COMPUTE_REGION=<your cluster's region>
    or
    CLOUDSDK_COMPUTE_ZONE=<your cluster's zone>

    # Name of GKE cluster
    CLOUDSDK_CONTAINER_CLUSTER=<your cluster's name>

    # (Optional) Project of GKE Cluster, only if you want helm to authenticate
    # to a GKE cluster in another project (requires IAM Service Accounts are properly setup)
    GCLOUD_PROJECT=<destination cluster's GCP project>

Setting the environment variables above will cause this step's `entrypoint` to
first run a command to fetch cluster credentials as follows.

    gcloud container clusters get-credentials --zone "$CLOUDSDK_COMPUTE_ZONE" "$CLOUDSDK_CONTAINER_CLUSTER"`

Then, `kubectl` and consequently `helm` will have the configuration needed to talk to your GKE cluster.

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud builds submit . --config=cloudbuild.yaml

## Using Helm

This builder supports two install options of Helm:
* The default one when the `tiller` gets installed into your GKE cluster (oh all those `tiller` security issues)
* `Tillerless Helm`, I wrote a [blog post](https://rimusz.net/tillerless-helm/) of using Helm local [tiller plugin](https://github.com/rimusz/helm-tiller) which solves all those `tiller` security issues, as `tiller` runs outside the GKE cluster.

Check the [examples](examples) folder for examples of using both Helm install options in `Cloud Build` pipelines.

You can test e.g. installing a chart via `Tillerless Helm`, running the following command.
**Note:** Do not forget to update `zone` and GKE `cluster` settings in the `cloudbuild.yaml` files.

    $ gcloud builds submit . --config=examples/chart-install-tillerless/cloudbuild.yaml

And to list Helm releases.

    $ gcloud builds submit . --config=examples/releases-list-tillerless/cloudbuild.yaml
