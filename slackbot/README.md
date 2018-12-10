# slackbot

Slackbot allows you to post build status messages to Slack.

![Slack screenshot](assets/screenshot.png)

## Getting started

If you are new to Google Cloud Build, we recommend you start by visiting the [manage resources page](https://console.cloud.google.com/cloud-resource-manager) in the Cloud Console, [enable billing](https://cloud.google.com/billing/docs/how-to/modify-project), [enable the Cloud Build API](https://console.cloud.google.com/flows/enableapi?apiid=cloudbuild.googleapis.com), and [install the Cloud SDK](https://cloud.google.com/sdk/docs/).

Clone this repository and build the builder:
```sh
gcloud builds submit . --config=cloudbuild.yaml
```

Follow instructions on the Slack website to [create a bot for your workspace](https://get.slack.help/hc/en-us/articles/115005265703-Create-a-bot-for-your-workspace).  Copy and paste the webhook URL, you'll need it in a moment.  You may also want to use the [Cloud Build logo](assets/cloud_build.png) as the bot's app icon, too.

## Generating notifications

Add the builder as the first step in your project's `cloudbuild.yaml`.  This triggers an independent "watcher" build which posts a status update whenever your main build completes - whether it's success or failure.

```yaml
steps:
- name: 'gcr.io/$PROJECT_ID/slackbot'
  args: [ '--build=$BUILD_ID',
          '--webhook=<your Slack webhook URL>' ]
...
```

## Examples

Examples showing both successful and unsuccessful builds are in the [examples](examples/) directory.
