[GCS cache]: https://cloud.google.com/cloud-build/docs/speeding-up-builds#caching_directories_with_google_cloud_storage
[Parallel Composite Uploads]: https://cloud.google.com/storage/docs/gsutil/commands/cp#parallel-composite-uploads

# Google Cloud Storage cache builder
To increase the speed of your build, you [can use a GCS bucket as a cache][GCS cache] with this builder.


## Building this builder
To build this builder, run the following command in this directory.

```bash
gcloud builds submit . --config=cloudbuild.yaml
```

## Using this builder
* Saving cache to the bucket with the `-s` flag
* Restoring cache from a bucket with the `-r` flag

| Option                | Description                                                         |
| --------------------- | ------------------------------------------------------------------- |
| `-r`                  | Mode: restore                                                       |
| `-s`                  | Mode: save                                                          |
| `-f file1 -f file2`   | File(s) to save, mandatory when saving                              |
| `-c gradle-cache.zip` | The compressed file name, optional, default `cache.zip`             |
| `-t 100`              | [Parallel Composite Uploads] threshold size, optional, default `50` |

The `GCS_BUCKET` environment variable must be passed to the build step in the `env` parameter.

Check the [examples](examples) folder for a common usage of this builder.
