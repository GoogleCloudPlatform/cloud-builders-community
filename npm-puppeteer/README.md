# Tool builder: `gcr.io/$PROJECT_ID/npm-puppeteer`

This Cloud Build builder runs the `npm` tool with Puppeteer and Chromium's dependancies installed for use with test scripts. 

It uses the [node:jessie-slim](https://hub.docker.com/_/node) as a base image. 

## Examples

See examples in the `examples` subdirectory.

## Building and publishing this builder

To build and publish this builder, run the following command in this directory.

```
gcloud builds submit . --config=cloudbuild.yaml
```
