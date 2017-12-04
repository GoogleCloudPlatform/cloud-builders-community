# boot

This build step invokes `boot` commands in [Google Cloud Container Builder](cloud.google.com/container-builder/).

Arguments passed to this builder will be passed to `boot` directly, allowing
callers to run [any boot
command](https://www.boot-clj.com).

## Credentials

You can securely pass credentials to `boot` [using encrypted
files](https://cloud.google.com/container-builder/docs/tutorials/using-encrypted-files).

See examples in the `examples` subdirectory.

## Status

This is unsupported demo-ware. Use at your own risk!
