# Google Cloud Container Builder community images

This repository contains source code for community-contributed builders used with the [Google Cloud Container
Builder API](https://cloud.google.com/container-builder/docs/).

These are not official Google products.

## How to use

First, [select or create a project](https://console.cloud.google.com/cloud-resource-manager), [enable billing](https://support.google.com/cloud/answer/6293499#enable-billing), [enable the Container Builder API](https://console.cloud.google.com/flows/enableapi?apiid=cloudbuild.googleapis.com), and [install and initialize the Cloud SDK](https://cloud.google.com/sdk/docs/).

Then, download the source code and build your builder.

### Example: using the `packer builder`

To use the `packer` builder on a Linux or Mac OS X workstation:

```sh
$ git clone http://github.com/GoogleCloudPlatform/cloud-builders-community
$ cd cloud-builders-community/packer
$ gcloud container builds submit --config cloudbuild.yaml .
```

View the builder image in Google Container Registry:

```sh
$ gcloud container images list --filter packer
```

Use the builder in your project's `cloudbuild.yaml`:

```yaml
- name: 'gcr.io/$PROJECT_ID/packer'
  args:
  - build
  - -var
  - project_id=$PROJECT_ID
  - packer.json
```

Examples of use are included in each tool's `examples` directory. [View a complete `packer` example](https://github.com/GoogleCloudPlatform/cloud-builders-community/tree/master/packer/examples/gce).

## Contributing

We welcome your contributions enthusiastically.  See [CONTRIBUTING](CONTRIBUTING.md) for more information on how to get started.  Please include a `cloudbuild.yaml` and at least one working example in your [pull request](https://help.github.com/articles/about-pull-requests/).

## License

This source code is licensed under Apache 2.0. Full license text is available in [LICENSE](LICENSE).

File issues here or e-mail `gcr-contact@google.com` if you have questions about
the usage of these build steps or the Cloud Container Builder API in general.

