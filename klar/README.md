# Klar
This builder allows you to use Klar to integrate Clair and Docker Registry (supports both Clair API v1 and v3).

To use this builder, you should configure a minimum number of variables as shown in the [example](./examples/cloudbuild.yaml). It is recommended that you create a new [service account](https://cloud.google.com/cloud-build/docs/how-to/service-account-permissions) with only `viewer` role and use its access token value for `DOCKER_PASSWORD`.

## Getting access token value of a service account
- Assuming you've created a service account with name `gcr-viewer` and downloaded the key for it, run the command below to activate the service account

    ```gcloud auth activate-service-account gcr-viewer --key-file=<path to json key file>```
- Then, print the access token for it by running the command:

    ```gcloud auth print-access-token gcr-viewer@<PROJECT-ID>.iam.gserviceaccount.com```

## Building this builder
Run the command below to build this builder

```
gcloud builds submit . --config=cloudbuild.yaml
```
