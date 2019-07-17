# [Terraform](https://www.terraform.io/docs) cloud builder

## Terraform cloud builder
This builder can be used to run the terraform tool in the GCE. From the Hashicorp Terraform [product page](https://www.terraform.io/):

> HashiCorp Terraform enables you to safely and predictably create, change, and improve infrastructure. It is an open source
> tool that codifies APIs into declarative configuration files that can be shared amongst team members, treated as code,
> edited, reviewed, and versioned.

### Building this builder
To build this builder, run the following command in this directory.
```sh
$ gcloud builds submit --config=cloudbuild.yaml
```

## Using this builder

### Terraform backend
Terraform stores state information about infrastructure it has provisioned. It uses this to plan out the delta between what your .tf files specifiy, and what's actually out there. This state can be stored in different ways by Terraform; it is configured via [backends](https://www.terraform.io/docs/backends/).

The default backend for Terraform is local, which will store state information the working directory in ```$ ./.terraform```. Most build platforms (including GCE) do not persist the working directory between builds. Losing this state information is no bueno.

There are a couple of options for managing Terraform state across builds:

###### Ignore the issue
In your build, you'll want to initialize terraform and refresh the local state. This is really not a good idea; it'll be slow and not multi-run safe (if multiple runs kick off concurrently, there'll be nastiness such as race conditions). [local_backend](examples/local_backend/README.markdown) is an example of this approach.
###### Persist the state in a GCS bucket manually
In your build, set up steps to manually fetch the state before running Terraform, then push it back up after Terraform is done. This will help by removing the need to init & refresh on every build; but will not address the concurrency issues.
###### Use a backend for remote storage
This is probably what you want to do. You'll still need to set up your GCS storage, and you'll need to configure the backend in your tf configurations. Some backends (happily, the [GCS](https://www.terraform.io/docs/backends/types/gcs.html) one does!) support locking of the remote state. This helps address the concurrency issue. [gcs_backend](examples/gcs_backend/README.markdown) is an example of this approach.

### Using this builder with Google Container Engine
To use this builder, your [builder service account](https://cloud.google.com/container-builder/docs/how-to/service-account-permissions) will need IAM permissions sufficient for the operations you want to perform. Adding the 'Kubernetes Engine Service Agent' role is sufficient for running the examples. Refer to the Google Cloud Platform [IAM integration page](https://cloud.google.com/container-engine/docs/iam-integration) for more info.

The article [Managing GCP projects with terraform](https://cloud.google.com/community/tutorials/managing-gcp-projects-with-terraform) gives a good strategy for administering projects in GCP with Terraform. If you intend to go beyond the examples, strongly consider an approach that isolates service accounts by function. A service account that can do 'all the things' is risky.

### Using this builder image anywhere else
This image can be run on any Docker host, without GCE. Why would you want to do this? It'll let you run Terraform locally, with no environment dependencies other than a Docker host installation. You can use the [local Cloud Build](https://cloud.google.com/cloud-build/docs/build-debug-locally) for this; but if you're curious or have
weird / advanced requirements (for example, if you want to run Terraform as a build step on another platform like Travis or Circle CI, and don't want to use Cloud Build, it is an option).

You'll need to:
 1. Provide a service account key file
 2. Mount your project directory at '/workspace' when you run docker
 ```sh
docker run -it --rm -e GCLOUD_SERVICE_KEY=${GCLOUD_SERVICE_KEY} \
  --mount type=bind,source=$PWD,target=/workdir \
  -w="/workdir" \
  gcr.io/$PROJECT_ID/terraform <command>
```
