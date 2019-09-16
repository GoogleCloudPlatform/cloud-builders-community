# Cancelot

Cancelot (thank you Twitter for the name :D) allows you to cancel previous builds running for the same branch, 
when you are using VCS triggered builds.

# Purpose

Cancelot was built because there is no out of the box solution by CloudBuild, in order to cancel a previous running 
build upon a new commit in the same branch. This can save a lot of build minutes that would be otherwise billed to the 
account.

# Deploying Cancelot

* Make any changes you need
* Navigate to Cancelot's folder and execute the following: `gcloud builds submit . --config=cloudbuild.yaml`
* Enjoy

## Using Cancelot

Add the builder as the first step in your project's `cloudbuild.yaml`:

```yaml
steps:
- name: 'gcr.io/$PROJECT_ID/cancelot'
  args: [ 
    '--current_build_id', '$BUILD_ID',
    '--branch_name', '$BRANCH_NAME'
  ]
```

Cancelot will be invoked when your build starts and will try to find any running jobs that match the following filter:

```
build_id != "[CURRENT_BUILD_ID]" AND 
source.repo_source.branch_name = "[BRANCH_NAME]" AND 
status = "WORKING" AND 
start_time<"[CURRENT_BUILD_START_TIME]"
```

After successfully fetching the list with the running builds that match the defined criteria, it loops and cancels 
each one.

### Contributing

After making any changes to Cancelot, please navigate to `test` folder & deploy the `cloudbuild.yaml`, like this:

```bash
gcloud builds submit . --config=cloudbuild.yaml --substitutions=BRANCH_NAME="test"
```

## Inspiration

Cancelot is heavily inspired by `slackbot` from CloudBuilders community
