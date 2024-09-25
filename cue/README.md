# CUE

[CUE](https://cuelang.org/) is a data constraint language for effective management of configurations, schemas and data.

This build step invokes `cue` commands in [Google Cloud Build](cloud.google.com/cloud-build/).

Arguments passed to this builder will be passed to `cue` directly, allowing
callers to run [any CUE CLI command](https://cuetorials.com/overview/cli-commands/).

## Building this builder

To build this builder, run the following command in this directory.

    gcloud builds submit . --config=cloudbuild.yaml
