# restore_cache

Restores the cache from the archive file defined by the provided `key`, either from the local `src` directory, or the remote GCS `bucket`

## Using this builder

All options use the form `--option=value` or `-o=value` so that they look nice in Yaml files.

| Option       | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| -b, --bucket | The cloud storage bucket to download the cache from. [optional]  |
| -s, --src    | The local directory in which the cache is stored. [optional] |
| -k, --key    | The cache key used for this cache file. [optional]           |

One of `--bucket` or `--src` parameters are required.  If `--bucket` then the cache file will be downloaded from the provided GCS bucket path.  If `--src` then the cache file will be read from the directory specified on disk.

The key provided by `--key` is used to identify the cache file.

### `checksum` Helper

This builder includes a `checksum` helper script, which you can use to create a simple checksum of files in your project to use as a cache key.

To use it in the `--key` arguemnt, simply surround the command with `$()`:

```bash
--key=build-cache-$(checksum cloudbuild.yaml)-$(checksum build.gradle)
```

## Examples

The following examples demonstrate build requests that use this builder.

### Restore a cache from a GCS bucket

This `cloudbuild.yaml` restores the files from the compressed cache file identified by `key` on the cache bucket provided, if it exists.

```yaml
- name: 'gcr.io/$PROJECT_ID/restore_cache'
  args:
  - '--bucket=gs://$CACHE_BUCKET/'
  - '--key=resources-$( checksum cloudbuild.yaml )'
```

### Restore a cache from a local file

This `cloudbuild.yaml` restores the files from the compressed cache file identified by `key` on the local filesystem, if it exists.

```yaml
- name: 'gcr.io/$PROJECT_ID/restore_cache'
  args:
  - '--src=/cache/'
  - '--key=resources-$( checksum cloudbuild.yaml )'
  volumes:
  - name: 'cache'
    path: '/cache'
```