This build step invokes the [Rocker builder](https://github.com/grammarly/rocker), an enhancement to `docker build` that adds some additional features.

The Rocker [README](https://github.com/grammarly/rocker/blob/master/README.md) has instructions and examples or see the examples subdirectory here.

This build step can be built and pushed to your GCP repository by running

    gcloud container builds submit --config=cloudbuild.yaml .
    
in this directory.
