# Docker-compose

This build step invokes `docker-compose` commands in [Google Cloud Build](cloud.google.com/cloud-build/).

Arguments passed to this builder will be passed to `docker-compose` directly,
allowing callers to run [any docker-compose
command](https://docs.docker.com/compose/reference/overview/).

## Setup

To make this cloud builder available in your active Google Cloud project:

```bash
cd cloud-builders-community/docker-compose
gcloud builds submit --config=cloudbuild.yaml .
```

To override the version of `docker-compose` being built, set the `_DOCKER_COMPOSE_VERSION` substitution:

```bash
cd cloud-builders-community/docker-compose
gcloud builds submit --config=cloudbuild.yaml --substitutions=_DOCKER_COMPOSE_VERSION="1.24.0"
```

You can find a list of releases and their version numbers [here](https://github.com/docker/compose/releases).

## Examples

See provided [hello-world](./examples/hello-world/) example.
