# save_cache

Saves the provided cache files to a compresssed archive, optionally uploading it to the provided cloud storage bucket.

## Using this builder

All options use the form `--option=value` or `-o=value` so that they look nice in Yaml files.

| Option       | Description                                                 |
| ------------ | ----------------------------------------------------------- |
| -b, --bucket | The cloud storage bucket to upload the cache to. [optional] |
| -o, --out    | The output directory to write the cache to. [optional]      |
| -k, --key    | The cache key used for this cache file. [optional]          |
| -p, --path   | The files to store in the cache. Can be repeated.           |

One of `--bucket` or `--out` parameters are required.  If `--bucket` then the cache file will be uploaded to the provided GCS bucket path.  If `--out` then the cache file will be stored in the directory specified on disk.

The key provided by `--key` is used to identify the cache file. Any other cache files for the same key will be overwritten by this one.

The `--path` parameters can be repeated for as many folders as you'd like to cache.  When restored, they will retain folder structure on disk.

### `checksum` Helper

As apps develop, cache needs change. For instance when dependencies are removed from a project, or versions are updated, there is no need to cache the older versions of dependencies. Therefore it's recommended that you update the cache key when these changes occur.

This builder includes a `checksum` helper script, which you can use to create a simple checksum of files in your project to use as a cache key.

To use it in the `--key` arguemnt, simply surround the command with `$()`:

```bash
--key=build-cache-$(checksum build.gradle)-$(checksum dependencies.gradle)
```

To ensure you aren't paying for storage of obsolete cache files you can add an Object Lifecycle Rule to the cache bucket to delete object older than 30 days.

## Examples

The following examples demonstrate build requests that use this builder.

### Saving a cache with checksum to GCS bucket

This `cloudbuild.yaml` saves the files and folders in the `path` arguments to a cache file in the GCS bucket `gs://$CACHE_BUCKET/`.  In this example the key will be updated, resulting in a new cache, every time the `cloudbuild.yaml` build file is changed.

```yaml
- name: 'gcr.io/$PROJECT_ID/save_cache'
  args:
  - '--bucket=gs://$CACHE_BUCKET/'
  - '--key=resources-$( checksum cloudbuild.yaml )'
  - '--path=.cache/folder1'
  - '--path=.cache/folder2/subfolder3'
```
