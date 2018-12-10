# slackbot examples

Edit the `cloudbuild.yaml` files in this directory to include your Slack webhook URL:

```yaml
steps:
- name: 'gcr.io/$PROJECT_ID/slackbot'
  args: [ '--build', '$BUILD_ID',
          '--webhook', '<Add your webhook URL here>' ]
- name: 'gcr.io/cloud-builders/docker'
  args: [ 'build', '.', '-f', 'Dockerfile-success']
```
