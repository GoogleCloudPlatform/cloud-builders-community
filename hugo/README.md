# Hugo

[Hugo](https://gohugo.io/) is a popular open-source static site generator.

This build step invokes `hugo` commands in [Google Cloud Container Builder](cloud.google.com/container-builder/).

Arguments passed to this builder will be passed to `hugo` directly, allowing
callers to run [any Hugo
command](https://gohugo.io/commands/).

## Usage:

To generate a static web-site from its sources:

```
steps:
- name: gcr.io/$PROJECT_ID/hugo
```
