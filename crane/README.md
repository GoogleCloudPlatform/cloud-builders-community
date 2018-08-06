# Skaffold (alpha)

This Container Builder build step runs
[`crane`](https://github.com/google/go-containerregistry/blob/master/cmd/crane/doc/crane.md).
By default, `crane` has access to the Google Container Registry in which the
build is running.

## Example: Fetching the digest of the builder

This example runs `crane digest` in a `cloudbuild.yaml`. This will print the
Docker digest of a given image.

```
steps:
- name: gcr.io/$PROJECT_ID/crane
  args: ['digest', 'gcr.io/$PROJECT_ID/crane']
```

## Building this builder

To build this builder and push the resulting image to the Container Registry
in your project, run the following command in this directory:

    $ gcloud builds submit . --config=cloudbuild.yaml

