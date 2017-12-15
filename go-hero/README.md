# Tool builder: `gcr.io/cloud-builders/go-hero`

This Container Builder build step runs the `hero` tool.

[Hero templates for Go language](https://github.com/shiyanhui/hero)

### When to use this builder

The `gcr.io/cloud-builders/go-hero` build step should be used when you want to run
the `hero` tool directly on your source, similar to how a developer uses the `hero`
tool locally to compile Hero templates.

### Related: [`gcr.io/cloud-builders/golang-project`](../golang-project/README.md)

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud container builds submit . --config=cloudbuild.yaml
