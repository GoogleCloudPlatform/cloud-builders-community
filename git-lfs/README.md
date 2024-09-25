# Container for Git with Git LFS as a build step on gcloud

See [docs](https://cloud.google.com/cloud-build/docs/concepts/custom-build-steps)

This container is based on the official Cloud Build Git docker container, but adds [Git LFS](https://git-lfs.github.com/).

## Usage:

```
steps:
- name: gcr.io/$PROJECT_ID/git-lfs
  id: git-lfs-help
  args: ['lfs', '--help']
- name: gcr.io/$PROJECT_ID/git-lfs
  id: git-lfs-build
  args: ['lfs', 'pull']
```
