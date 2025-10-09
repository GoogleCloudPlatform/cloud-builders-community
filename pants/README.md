# Pants

The [pants](https://pantsbuild.org) build system is an open-source mutilingual build tool.


## Building this builder

gcloud builds submit . --config=cloudbuild.yaml --substitutions=_PY_VERSION="3.11"

## Using Pants

Assuming you built and pushed the pants builder for python 3.11, an example usage might
look like:

```
steps:
  - name: gcr.io/$PROJECT_ID/pants:3.11
    args:
      - package
      - path/to/my:package
```
