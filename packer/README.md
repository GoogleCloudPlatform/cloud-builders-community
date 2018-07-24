# Packer

This build step invokes `packer` commands in [Google Cloud Build](https://cloud.google.com/cloud-build).

Arguments passed to this builder will be passed to `packer` directly, allowing
callers to run [any Packer
command](https://www.packer.io/docs/commands/index.html).

## Credentials

You can securely pass credentials to `packer` [using encrypted
files](https://cloud.google.com/cloud-build/docs/tutorials/using-encrypted-files).

See examples in the `examples` subdirectory.

## Status

This is unsupported demo-ware. Use at your own risk!
