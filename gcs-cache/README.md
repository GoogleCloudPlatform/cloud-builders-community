# GCS Cache

This Container Builder build step enhances the
[`gsutil`](https://github.com/GoogleCloudPlatform/cloud-builders/tree/master/gsutil)
builder with `push.sh` and `pull.sh` scripts that ease use of a GCS cache
between builds.

## Building this builder

To build this builder, run the following command in this directory:

    $ gcloud container builds submit . --config=cloudbuild.yaml

## Using this builder
