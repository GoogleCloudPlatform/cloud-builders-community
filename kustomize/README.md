# Kustomize

This Container Builder build step runs the Kustomize
[`kustomize`](https://github.com/kubernetes-sigs/kustomize) tool.

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud builds submit . --config=cloudbuild.yaml
