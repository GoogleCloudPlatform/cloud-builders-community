# Godep

This build step provides a container based on the golang container that adds in
the [Godep](https://github.com/tools/godep) tool.  Godep is an official golang
tool for dependency management.

This image uses godep as the entrypoint and takes in go commands as arguments.
An example use case for Google Container Builder:

```
- name: gcr.io/cloud-builders-community/godep
  dir: "/workspace/src/github.com/GoogleCloudPlatform/k8s-stackdriver/event-exporter"
  args: ["go", "test", "./..."]
```

The above example wraps the go test command in godep, which fetches packages
from the Godeps directory to satisfy dependencies.
