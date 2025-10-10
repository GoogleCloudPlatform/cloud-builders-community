# Katalon in Google Cloud Build

This demonstrates using the
[`Katalon`](https://github.com/katalon-studio/docker-images) tool in Google Cloud
Build to produce Docker images in a registry, instead of using `docker build`
and `docker push`.

## Steps:

1. Modify the `cloudbuild.yaml` to define the steps to build your Docker image using Katalon.
2. Define the custom `Dockerfile` for Katalon, which installs necessary packages and defines an entry point.
3. Use the Cloud Build trigger to automate the process of building and pushing the Katalon image to your Container Registry, tagged with both `latest` and `9.6.0`.

## Trigger the build locally

To trigger the build locally using `gcloud`, navigate to your project root directory and run the following command:

```bash
gcloud builds submit --config=cloudbuild.yaml .
```

This command will submit your build to Cloud Build using the cloudbuild.yaml configuration file, which includes the steps for tagging the image with both latest and 9.6.0.

