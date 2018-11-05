# Tool builder: `gcr.io/$PROJECT_ID/npm-puppeteer`

This Container Builder build step runs the `npm` tool but with the necessary dependencies for [puppeteer](https://github.com/GoogleChrome/puppeteer).

It uses the small alpine-node base.

**Note:** This is based on [yarn-puppeteer](../yarn-puppeteer/README.md), but using npm instead.

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud builds submit . --config=cloudbuild.yaml
