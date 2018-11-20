# CFT GCE Build

This build creates a GCE instance by executing `cft` command.

Note: This example assumes that you have built the `cft` build step and pushed it to
`gcr.io/$PROJECT_ID/cft`.

## Executing the CFT Builder

To run this example, make sure you have assigned the [Deployment Manager
Editor](https://cloud.google.com/iam/docs/understanding-roles#deployment_manager_roles)
role to [your Cloud Build service
account](https://cloud.google.com/cloud-build/docs/securing-builds/set-service-account-permissions),
and run:
```
    gcloud builds submit . --config=cloudbuild.yaml
```