# Tool builder: `gcr.io/$PROJECT_ID/yarn-puppeteer`

This Container Builder build step runs the `yarn` tool but with the necessary dependencies for puppeteer.

It uses the small alpine-node base.

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud container builds submit . --config=cloudbuild.yaml
