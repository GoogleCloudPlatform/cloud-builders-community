# unzip

This is a tool build to simply invoke the
[`unzip`](https://linux.die.net/man/1/unzip) command.

Arguments passed to this builder will be passed to `unzip` directly.

## Examples

The following examples demonstrate build requests that use this builder.

### Extract an existing archive 'my-archive.jar' into a target directory '/contents'

```
steps:
- name: gcr.io/$PROJECT_ID/unzip
  args: ["-d", "contents", "my-archive.jar"]
```
