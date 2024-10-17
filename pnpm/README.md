# pnpm builder for Google Cloud Build

This image can be used in Google Cloud Run to have access to `pnpm`.

## Submit the image to your project

```bash
gcloud builds submit --tag gcr.io/<your-project-id>/pnpm-builder
```

## Example usage on a `cloudbuild.yaml`

```yaml
steps:
  - id: install
    name: gcr.io/$PROJECT_ID/pnpm-builder
    entrypoint: pnpm
    args: ['install', '--frozen-lockfile']
  - id: test
    name: gcr.io/$PROJECT_ID/pnpm-builder
    entrypoint: pnpm
    args: ['run', 'test']
```
