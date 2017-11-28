<<<<<<< HEAD
# Google Cloud Container Builder community images

This repository contains source code for community-contributed Docker images. You can use these images as build steps for
[Google Cloud Container Builder](https://cloud.google.com/container-builder/docs/).

These are not official Google products.

## How to use a community-contributed build step

Google Container Builder executes a build as a series of build steps. Each build step is run in a Docker container. See
the [Container Builder documentation](https://cloud.google.com/container-builder/docs/overview) for more details
about builds and build steps.

### Before you begin

1.  Select or create a [Google Cloud project](https://console.cloud.google.com/cloud-resource-manager).
2.  Enable [billing for your project](https://support.google.com/cloud/answer/6293499#enable-billing).
3.  Enable [the Container Builder API](https://console.cloud.google.com/flows/enableapi?apiid=cloudbuild.googleapis.com).
4.  Install and initialize [the Cloud SDK](https://cloud.google.com/sdk/docs/).

### Build the build step from source

To use a community-contributed Docker image as a build step, you need to download the source code from this
repo and build the image.

The example below shows how to download and build the image for the `packer` build step on a Linux or Mac OS X workstation:

1. Clone the `cloud-builders-community` repo:

   ```sh
   $ git clone http://github.com/GoogleCloudPlatform/cloud-builders-community
   ```

2. Go to the directory that has the source code for the `packer` Docker image:

   ```sh
   $ cd cloud-builders-community/packer
   ```

3. Build the Docker image:

   ```
   $ gcloud container builds submit --config cloudbuild.yaml .
   ```

4. View the image in Google Container Registry:

   ```sh
   $ gcloud container images list --filter packer
   ```

### Use the build step with Container Builder build

Once you've built the Docker image, you can use it as a build step in a Container Builder build.

For example, below is the `packer` build step in a YAML
[config file](https://cloud.google.com/container-builder/docs/build-config), ready to be used in a Container Builder build:

   ```yaml
   - name: 'gcr.io/$PROJECT_ID/packer'
     args:
     - build
     - -var
     - project_id=$PROJECT_ID
     - packer.json
   ```

Each build step's `examples` directory has an example of how you can use the build step. See the
[example for the `packer` builder](https://github.com/GoogleCloudPlatform/cloud-builders-community/tree/master/packer/examples/gce).

## Contributing

We welcome contributions!  See [CONTRIBUTING](CONTRIBUTING.md) for more information on how to get started.
Please include a `cloudbuild.yaml` and at least one working example in your
[pull request](https://help.github.com/articles/about-pull-requests/).

## License

This source code is licensed under Apache 2.0. Full license text is available in [LICENSE](LICENSE).

## Support

*   File issues here in GitHub about the community build steps.
*   Email `gcr-contact@google.com` if you have questions about build steps or Google Cloud Container Builder.

=======
# Container Builder Remote Build Step

## THIS IS NOT AN OFFICIAL GOOGLE PRODUCT

## Intro

### What?

Some continuous integration workloads require special builder types. You may
require things like:

1. High CPU/Memory machines
1. Custom image
1. GPUs attached
1. Fast or large disks
1. Machines in a particular network
1. Pre-emptibility

In these cases you can leverage Container Builder to trigger your builds and
manage their workflow but run the actual build steps on an instance with
exactly the configuration you need.

### How?

When using the remote-builder image, the following will happen:

1. A temporary SSH key will be created in your Container Builder workspace
1. A instance will be launched with your configured flags
1. The workpace will be copied to the remote instance
1. Your command will be run inside that instance's workspace
1. The workspace will be copied back to your Container Builder workspace

## Usage

In order to use this step you can configure your build step as follows:

```
steps:
- name: gcr.io/cloud-solutions-images/remote-builder:v0.3.0
  env:
    - COMMAND=ls -la
```

This will launch an instance with the default parameters and then run the
command `ls -la` inside the instance's workspace.

## Configuration

The following options are configurable via environment variables passed to the
build step in the `env` parameter:

| Options       | Description   | Default |
| ------------- | ------------- | ------- |
| USERNAME  | Username to use when logging into the instance via SSH  | `cloud-user` |
| REMOTE_WORKSPACE  | Location on remote host to use as workspace | `/home/${USERNAME}/workspace/` |
| INSTANCE_NAME  | Name of the instance that is launched  | `builder-$UUID` |
| ZONE  | Compute zone to launch the instance in | `us-central1-f` |
| INSTANCE_ARGS| Parameters to the instance creation command. For a full list run `gcloud compute instances create --help`| `--preemptible` |

## Quick Start

In the following example, you will run a script inside of containers on two instances in
parallel. You will use the Container Optimized OS image to provide an image with Docker
pre-installed. The build request runs the `test/no-op.sh` script from this directory.

1. Create the build request:

```shell
cat > cloudbuild.yaml <<EOF
steps:
- name: gcr.io/cloud-solutions-images/remote-builder:v0.3.0
  waitFor: ["-"]
  env:
    # Use Container Optimized OS
    # https://cloud.google.com/container-optimized-os/docs/
    - INSTANCE_ARGS=--image-project cos-cloud --image-family cos-stable
    - USERNAME=cloud-user
    # Run a script from the local build context in a Docker container
    - COMMAND=docker run -v /home/cloud-user/workspace:/workspace ubuntu:16.04 bash -xe /workspace/test-scripts/no-op.sh
- name: gcr.io/cloud-solutions-images/remote-builder:v0.3.0
  waitFor: ["-"]
  env:
    # Use Container Optimized OS
    # https://cloud.google.com/container-optimized-os/docs/
    - INSTANCE_ARGS=--image-project cos-cloud --image-family cos-stable
    - USERNAME=cloud-user
    # Run a script from the local build context in a Docker container
    - COMMAND=docker run -v /home/cloud-user/workspace:/workspace ubuntu:16.04 bash -xe /workspace/test-scripts/no-op.sh
EOF
```

2. Run submit the build with the local context as the workspace. 

```shell
gcloud container builds submit --config cloudbuild.yaml .
```

3. You should now see 2 instances being provisioned in parallel then the `test/no-op.sh` being
run inside containers based on the `ubuntu:16.04` Docker image.
>>>>>>> d4446d0... Initial Commit
