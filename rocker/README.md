This build step invokes the Rocker builder https://github.com/grammarly/rocker, an enhancement to `docker build` that adds some additional features.

See the Rocker documentation for extensive examples.

This build step can be built and pushed to your GCP repository by running

    gcloud container builds submit --config=cloudbuild.yaml .
    
in this directory.
