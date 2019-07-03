# Hello World

## Development Build, Test & Deploy

```
cloud-build-local --dryrun=false \
  --substitutions=_DEPLOYMENT_NAME=hello-world-development \
  .
```

Equivalently, a trigger (perhaps responding to code pushed to any branch other than `master`) can be created on a repository containing this project, which makes the appropriate substitution.

## Staging Build, Test & Deploy

```
cloud-build-local --dryrun=false \
  --substitutions=_DEPLOYMENT_NAME=hello-world-staging \
  .
```

Equivalently, a trigger (perhaps responding to code pushed to `master`) can be created on a repository containing this project, which makes the appropriate substitution.

## Production Build, Test & Deploy

```
cloud-build-local --dryrun=false \
  --substitutions=_DEPLOYMENT_NAME=hello-world \
  .
```

Equivalently, a trigger (perhaps responding to a tag) can be created on a repository containing this project, which makes the appropriate substitution.
