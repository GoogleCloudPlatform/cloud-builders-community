# Packer

This build step invokes `packer` commands in [Google Cloud Build][cloud-build].

Arguments passed to this builder will be passed to [`packer`][packer] directly, allowing callers to
run [any Packer command][packer-commands].

[cloud-build]: https://cloud.google.com/cloud-build

[packer]: https://www.packer.io

[packer-commands]: https://www.packer.io/docs/commands

## Building this Builder

Before using this builder in a Cloud Build config, it must be built and pushed to the registry in
your project. Run the following command in this directory:

```
gcloud builds submit .
```

> **Advanced builder building:** To specify a particular version of packer, provide the packer version
> number, and the checksum of that version's zip archive, as Cloud Build [substitutions][substitutions]:
> ```
> gcloud builds submit --substitutions=_PACKER_VERSION=1.7.8,_PACKER_VERSION_SHA256SUM=8a94b84542d21b8785847f4cccc8a6da4c7be5e16d4b1a2d0a5f7ec5532faec0 .
> ```

[substitutions]: https://cloud.google.com/cloud-build/docs/configuring-builds/substitute-variable-values#using_user-defined_substitutions

## Credentials

You can securely pass credentials to `packer` [using encrypted files][cloud-build-encrypted-files].

See examples in the `examples` subdirectory.

[cloud-build-encrypted-files]: https://cloud.google.com/cloud-build/docs/tutorials/using-encrypted-files

## Status

This is unsupported demo-ware. Use at your own risk!
