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
> gcloud builds submit --substitutions=_PACKER_VERSION=1.8.3,_PACKER_VERSION_SHA256SUM=0587f7815ed79589cd9c2b754c82115731c8d0b8fd3b746fe40055d969facba5 .
> ```

[substitutions]: https://cloud.google.com/cloud-build/docs/configuring-builds/substitute-variable-values#using_user-defined_substitutions

## Credentials

You can securely pass credentials to `packer` [using encrypted files][cloud-build-encrypted-files].

See examples in the `examples` subdirectory.

[cloud-build-encrypted-files]: https://cloud.google.com/cloud-build/docs/tutorials/using-encrypted-files

## Status

This is unsupported demo-ware. Use at your own risk!
