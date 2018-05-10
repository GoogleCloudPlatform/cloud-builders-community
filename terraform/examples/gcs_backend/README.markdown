# GCS back end example

This is an example of how to use the Terraform cloud builder, using a GCS backend.

### Building this builder
To build this builder, run the following command in this directory.
```sh
$ gcloud container builds submit . --config=cloudbuild.yaml
```

## Using this builder

1. Create a new GCloud project.
1. Open the [GCP IAM console](https://console.cloud.google.com/iam-admin) and select your GCloud project
1. Find the 'Cloud Container Builder' service account for your project. It will
   - Have the 'Cloud Container Builder' role
   - The Member ID will be <project ID>@cloudbuild.gserviceaccount.com
1. Edit the permission for that service account and add the 'Kubernetes Engine Service Agent' role.
1. Clone this project
1. [Build the Terraform cloud builder](../../README.markdown)
1. Navigate to this directory
1. Build this builder

## What's it do?
This builder will create a cluster named 'terraform-builder-gcs-backend' in your project, per main.tf. It will then destroy it. All told this will take around 5 minutes.

This builder will create a bucket to hold data for the Terraform GCS back end. [cloubuild.yaml](cloudbuild.yaml) does this in the first step. The bucket name is configured in cloudbuild.yml on line 26. After this bucket is created, you'll need to remove the first step; it will fail because the bucket already exists.

The bucket will persist until you delete it. Don't forget to do that, once you're done with the examples! You can browse GCS storage [here](https://console.cloud.google.com/storage/browser).

## Parameterization
It's worth noting how Terraform passes variables via the command line. In cloudbuild.yaml, the project name is passed into the build steps as an environment variable
```yaml
"TF_VAR_project-name=${PROJECT_ID}"
```
TF_VAR is a prefix Terraform uses to identify tf variables; the rest maps to a variable defined in variables.tf. Depending on your needs, you may not want to parameterize values like this; you could just hard-code them in the tf files.

