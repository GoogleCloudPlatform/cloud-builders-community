# lerna

> Lerna is a tool that optimizes the workflow around managing multi-package repositories with git and npm.

Arguments passed to this builder will be passed to `lerna` directly.

The node image used to build the lerna builder and the lerna package to install
can be overridden by substitutions. Please see
[cloudbuild.yaml](./cloudbuild.yaml) for available substitutions and their
default values.

To build:

```sh
gcloud builds submit \
  --config=cloudbuild.yaml \
  # optional: override the node image, tag and lerna package
  --substitutions=_NODE_IMAGE=node,_NODE_TAG=12,LERNA_PACKAGE=lerna \
  # optional: build in a specific project
  --project=<GCR_PROJECT>
```
