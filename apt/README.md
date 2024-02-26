# apt

The `gcr.io/$PROJECT_ID/apt` image is maintained by the Cloud Build team,
but it may not support the most recent versions of commands required to build
apt packages.

-------

## Examples

The following examples demonstrate build requests that use `apt`.

### Build apt packages

This `cloudbuild.yaml` simply invokes `fakeroot dpkg-deb` to build apt package.

```yaml
steps:
- name: 'gcr.io/$PROJECT_ID/apt'
  script: |
    fakeroot dpkg-deb --build package-root
```
