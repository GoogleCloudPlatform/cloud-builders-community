# Setup

To make this cloud builder available in your active Google Cloud project:
```bash
cd cloud-builders-community/docker-compose
gcloud builds submit --config=cloudbuild.yaml .
```

The build defaults to using the latest version of `docker-compose` at the time of writing: `1.25.5`. To use a different version:
```bash
cd cloud-builders-community/docker-compose
gcloud builds submit --config=cloudbuild.yaml --substitutions=_DOCKER_COMPOSE_VERSION="1.24.0"
```

You can find a list of releases and their version numbers [here](https://github.com/docker/compose/releases).


# Example Build
The provided `cloudbuild.yaml` simply invokes `docker-compose up` for a service that says "Hello world".

```bash
cd cloud-builders-community/docker-compose/examples/hello-world
gcloud builds submit --config=cloudbuild.yaml .
```
