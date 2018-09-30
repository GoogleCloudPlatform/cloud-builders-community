# Packer GCE build

This directory contains an example that creates a GCE image using `packer`.

Note: This example assumes that you have built the `packer` build step and pushed it to
`gcr.io/$PROJECT_ID/packer`.

## Preparing GCE Credentials for the Packer Builder

To provide proper credentials to the `packer` builder:

1.  Create a [Service
    Account](https://cloud.google.com/iam/docs/service-accounts) with `Editor`
    access to your GCP project.
1.  Create a [Service Account
    Key](https://cloud.google.com/iam/docs/creating-managing-service-account-keys)
    credentials file in `json` format.
1.  Encrypt the `json` credentials file into `secrets.enc` as described in the
    [using encrypted
    files](https://cloud.google.com/cloud-build/docs/tutorials/using-encrypted-files)
    tutorial for Cloud Build. Note that this `packer` example assumes that
    you have encrypted using a key named `key` in keyring `packer`; this is the
    key used to decrypt in `cloudbuild.yaml`.

## Executing the Packer Builder

Now that you've got `secrets.enc` in place, execute your `packer` build:

    gcloud builds submit --config=cloudbuild.yaml .
