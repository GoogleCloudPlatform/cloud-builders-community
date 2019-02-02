# Codecov

This build step invokes `codecov` commands in [Google Cloud Build](https://cloud.google.com/cloud-build).

Arguments passed to this builder will be passed to `codecov` directly.

Example: 
```
- name: 'gcr.io/node-example-gke/codecov'
  args: 
  - '-t'
  - 'my-token'
  - '-C'
  - '$COMMIT_SHA'
  - '-B'
  - '$BRANCH_NAME'
  - '-b'
  - '$BUILD_ID'
  - '-T'
  - '$TAG_NAME'
```
