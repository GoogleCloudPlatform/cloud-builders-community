# GCS Syncer

## About

This builder is a tool to perform exact sync with/from Google Cloud Storage (GCS) periodically. `gcs-syncer` is useful to sync data between two different gcs buckets/gcs paths or between gcs bucket to local path. Personally, I used `gcs-syncer` mainly as a sidecar container in my kubernetes pod/deployments so that I can achieve separation of concern between my app and the data sync logic.

This project is inspired from my SO question here: https://stackoverflow.com/questions/54877057/how-to-separate-application-and-data-syncing-implementations-in-kubernetes.

## 1. Building gcs-syncer

Build this image and add it to your GCS repo

```
git clone git@github.com:GoogleCloudPlatform/cloud-builders-community
cd cloud-builders/gcs-syncer
gcloud builds submit . --config=cloudbuild.yaml
```

## 2. Prerequisites for running gcs-syncer

In order to run `gcs-syncer`, you'll need service account, generate a `GOOGLE_APPLICATION_CREDENTIALS` key and grant the required IAM permissions to the bucket for your service accounts.

1. To create a service acccount, you can follow the guide in https://cloud.google.com/iam/docs/creating-managing-service-accounts#iam-service-accounts-list-console

2. To generate a service account key for existing accounts, you can follow the guide in https://cloud.google.com/iam/docs/creating-managing-service-account-keys#iam-service-account-keys-create-console

3. The required IAM permissions that you should grant to an account in order to be able to perform `rsync -d` is available in  this page: https://cloud.google.com/storage/docs/access-control/iam-gsutil. To grant IAM permissions to your service account, you can follow the guide in https://cloud.google.com/iam/docs/granting-roles-to-service-accounts

## 3. Using gcs-syncer

Arguments passed to `gcs-syncer` will be passed to `sync.sh` directly, allowing callers to run `gsutil rsync` operations
with any source and destination paths.

When `gcs-syncer` is run, it tries to use the default gcloud credentials of the caller. You can (and should in the case of docker container) override this with service account credentials by setting the environment variable `GOOGLE_APPLICATION_CREDENTIALS` of the image to point to a service account credentials path.

When `gcs-syncer` is run, it will run indefinitely until it receives user signal to stop, which depends on how you run `gcs-syncer`.

There are three ways of running `gcs-syncer`. Before running `gcs-syncer` for docker / kubernetes, ensure that:

1. you've built the gcs-syncer image first using (step 1)
2. you've generated a service account key (step 2)


### Running gcs-syncer locally as a bash script
```bash
# Periodically sync gcs to local data every 5 seconds with default gcloud credentials
./sync.sh gcr.io/[YOUR_PROJECT_ID]/gcs-syncer gs://bucket-name/path/to/data/ data 5

GOOGLE_APPLICATION_CREDENTIALS=[Path to credential files]
# Periodically sync between gcs bucket every 5 seconds with a service account credentials
./sync.sh gcr.io/[YOUR_PROJECT_ID]/gcs-syncer gs://bucket-name/path/to/data/ gs://another-bucket-name/path/to/data/ 5
```

To stop the running docker container, you can send kill signals to the process ( hit `ctrl + c` or run `kill -9 [PID_OF_SYNC.SH]`)

Ref: ds.nyu.edu/academics/ms-in-data-science/

### Running gcs-syncer locally as docker image

```bash
# Pull your latest image to local
docker pull gcr.io/[YOUR_PROJECT_ID]/gcs-syncer

# Sync from gcs to local data path every 5 seconds
docker run --rm --name gcs-syncer-local  \
    -v [PATH_TO_ACCOUNT_CREDENTIALS]:/var/key.json \
    -v `pwd`/data:/data \
    -e GOOGLE_APPLICATION_CREDENTIALS=/var/key.json \
    gcr.io/[YOUR_PROJECT_ID]/gcs-syncer gs://bucket-name/path/to/data/ data 5
```

To stop the running docker container, you can run `docker stop gcs-syncer-local`.

### Running gcs-syncer as cloud builder

See `examples/cloudbuild.yaml`

Note that when executed from Cloud Build environment, commands are executed with credentials of the [builder service account](https://cloud.google.com/cloud-build/docs/permissions) of the project. You can use `GOOGLE_APPLICATION_CREDENTIALS` to override this.

When using cloud builder to run `gcs-syncer`, you can stop the running process by visiting [google cloud console for cloud builder](https://console.cloud.google.com/cloud-build/builds) and cancelling the running build process.

### Running gcs-syncer as a sidecar container

See `examples/deployment.yaml` on how you can use `gcs-syncer` with kubernetes.

You'll need to first create a kubernetes secret of your secret credentials and mount this secret in your kubernetes pods.

```bash
# 1. Install kubernetes if you haven't installed it yet
gcloud components install kubectl

# 2. Have your secret credential key file ready with the correct IAM access to the buckets, then create a kubernetes secret:
kubectl create secret generic {{ your-credentials-key }} --from-file=credentials.json=[ PATH-TO-KEY-FILE.json ]

# 3. Apply the deployment
kubectl apply -f examples/deployment.yaml
```
