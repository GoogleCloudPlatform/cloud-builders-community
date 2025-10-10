# Poetry

See [docs](https://python-poetry.org/docs/cli/)

This cloud-builder offer the [`poetry`(https://python-poetry.org) command to be used by build and publish packages, for example.


## Usage

Your arguments passed to this builder will be send to `poetry` directly. For example:

```yaml
---
steps:
  - name: 'gcr.io/${PROJECT_ID}/poetry'
    id: 'publish package'
    args: ['publish', '-r', 'pypi', '--build']
```

