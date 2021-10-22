# CloudBuild with Terraform at Scale

CloudBuild gives your Terraform deployment boost and scale allowing you to concurrently deploy your infrastructure across dozens of GCP regions with a single command. This is how you do it:

1. Ensure that $PROJECT_ID environment variable is set:

    ```bash
    export PROJECT_ID = [PROJECT_ID]
    ```

1. Enable CloudBuild on your project:

    ```bash
    gcloud services enable cloudbuild.googleapis.com container.googleapis.com compute.googleapis.com cloudresourcemanager.googleapis.com
    ```

1. Build the Terraform builder

    ```bash
    # run from the terraform directory (cd ../..)
    gcloud builds submit --config=cloudbuild.yaml
    ```

1. Give permissions to the Cloud Build service account to create infrastructure. This example uses `roles/owner` for simplicity, but you should restrict to the requried roles.

    ```bash
    export PROJECT_NUM=$(gcloud projects describe $PROJECT_ID --format="value(projectNumber)")
    gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member serviceAccount:$PROJECT_NUM@cloudbuild.gserviceaccount.com \
    --role roles/owner
    ```
1. Increase your build and operation GetRequests: https://pantheon.corp.google.com/apis/api/cloudbuild.googleapis.com/quotas?project=alexbu-test-20211022
1. Run multi-build