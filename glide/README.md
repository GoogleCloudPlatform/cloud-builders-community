# glide

This Cloud Build build step runs the [`glide`](https://glide.sh) tool.
[`glide`](https://glide.sh) is used to manage Go source dependencies.

### When to use this builder

This build step should be used when you want to run
the [`glide`](https://glide.sh) tool on your source.

`gcr.io/cloud-builders/go` is used to build a go workspace.

## Examples

-   [Build glide](examples/build-glide) is a basic example that clones the
    [`glide`](https://glide.sh) source from its
    [GitHub repository](https://github.com/Masterminds/glide), uses `glide
    install` to pull in the correct versions of all dependencies, and the
    [`go`](https://github.com/GoogleCloudPlatform/cloud-builders/blob/master/go/README.md)
    builder to build the `glide` binary.

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud container builds submit . --config=cloudbuild.yaml
