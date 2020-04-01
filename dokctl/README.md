## Dokctl - Digital Ocean Kubectl

This will help digital ocean kubernetes users for their CI/CD integration with Google CloudBuild. 

## Using this builder with Digital Ocean Kubernetes Cluster

Example cloudbuild.yaml is available in examples/pods-list directory

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud builds submit . --config=cloudbuild.yaml
