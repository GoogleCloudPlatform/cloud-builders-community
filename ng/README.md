# The ng (Angular CLI) Cloud Builders

This build step invokes the `ng` command provided by the [Angular CLI](https://github.com/angular/angular-cli) distributed in [Google Cloud Build](cloud.google.com/cloud-build/).

Arguments passed to this builder will be passed to the `ng` command directly,
allowing callers to run [any `ng`
command](https://github.com/angular/angular-cli/wiki#additional-commands/).

## Available builders

For convenience, we have included different versions of the Angular CLI:
- `gcr.io/cloud-builders/ng:v1`: provides the v1.* legacy branch
- `gcr.io/cloud-builders/ng:v6`: provides the v6.* branch
- `gcr.io/cloud-builders/ng:latest`: provides the latest stable branch
- `gcr.io/cloud-builders/ng`: same as latest
- `gcr.io/cloud-builders/ng:next`: provides the next unstable branch

## How to use?

In order to use call one of these builder, simply invoke the builder (and version), for instance:

```
steps:
- name: 'gcr.io/$PROJECT_ID/ng:next'
  args: ['build', '--prod']
```

Or, if you are maintaining a legacy Angular project:
```
steps:
- name: 'gcr.io/$PROJECT_ID/ng:v1'
  args: ['test', '--sourcemap=false']
```

See the `examples` folder for a complete example.


## Building these builders

To build these builders, run the following command in this directory:

    $ gcloud builds submit . --config=cloudbuild.yaml
