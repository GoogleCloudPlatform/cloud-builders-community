# Swift example

This `cloudbuild.yaml` invokes a simple `swift --version` command:
```
gcloud container builds submit --config=cloudbuild.yaml .
```
