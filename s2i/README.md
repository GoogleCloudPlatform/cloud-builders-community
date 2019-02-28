# s2i builder

This bulid step invokes `s2i` commands in [Google Cloud Build](cloud.google.com/cloud-build/)

[s2i](https://github.com/openshift/source-to-image) is a tookit and workflow for building reproducible Docker images from source code.

Arguments passed to this builder will be passed to `s2i` directly, allowing callers to build images with `s2i` directly.

## Examples

See examples in the `examples` subdirectry.
