# Swift

This build step invokes `swift` commands in [Google Cloud Container Builder](https://cloud.google.com/container-builder/).

Arguments passed to this builder will be passed to `swift` directly,
allowing callers to run [any swift command](https://swift.org/lldb/).

## Swift Versions
This build step supports multiple swift versions depending on your needs.

### Available swift versions
- `gcr.io/cloud-builders/swift:bionic` : The latest release on Ubuntu 18.04 (bionic)
- `gcr.io/cloud-builders/swift:xenial` : The latest release on Ubuntu 16.04 (xenial)
- `gcr.io/cloud-builders/swift:latest-dev` : The latest development snapshot for the upcoming release.
- `gcr.io/cloud-builders/swift:4.2` :  The 4.2.2 release branch
- `gcr.io/cloud-builders/swift:4.2.4` :  The 4.2.4 release branch 
- `gcr.io/cloud-builders/swift:5.0-xenial` :  The 5.0.2 release branch on Ubuntu 16.04 (xenial)
- `gcr.io/cloud-builders/swift:5.0-bionic` :  The 5.0.2 release branch on Ubuntu 18.04 (bionic)
