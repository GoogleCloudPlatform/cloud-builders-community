# Harness Chaos

## Introduction
The harness-chaos builder step is used to launch chaos experiments and test the resilience of the applications. This tool helps users easily simulate various failure scenarios and uncover potential vulnerabilities and weaknesses in the system/applications.This can be integrated with the CD pipelines to test the resiliency of the application before doing the releases

## Prerequisites

**Identifiers(Account ID, Organization ID, Project ID)**

To perform any operation in Harness Chaos Engineering(HCE) platform, users need to get onboarded and have the identifiers. Here are the references to get started:
- [User Onboarding Guide](https://developer.harness.io/docs/platform/get-started/onboarding-guide) 
- [Create organizations and projects](https://developer.harness.io/docs/platform/organizations-and-projects/create-an-organization)

**Create API Key**
To run any operation, users will need an API Key to authorize the same. API Key needs to be created and passed as flag in the command demonstrated below. Here is the document to [create and manage API keys](https://developer.harness.io/docs/platform/automation/api/add-and-manage-api-keys/).

**Create a Chaos Experiment**

Next step is to create a chaos experiment and get the experiment ID to launch the chaos experiment using cloud builder. Once the onboarding is complete, users can go to the Chaos Engineering Module and follow the steps mentioned in documentation to [create a chaos experiment](https://developer.harness.io/docs/chaos-engineering/get-started/tutorials/chaos-experiment-from-blank-canvas).

## How to run launch chaos experiment using cloud builder

Once the chaos experiment is created and we have the values for the required flags, here’s an example demonstrating how to launch and validate a chaos experiment using cloud builder. Users can pass the `--experiment-id` flag along with the `--expected-res-score` flag to ensure that the actual resilience score of the experiment run meets the expected threshold, along with other necessary flags.
We recommend to store and use the API-KEY as a secret.

**Example:**
```
steps:
  - name: 'us-central1-docker.pkg.dev/$PROJECT_ID/harness-chaos/harness-chaos'
    secretEnv: ['API_KEY']
    args: ['experiment', 'run', '--account-id=${_ACCOUNT_ID}','--org-id=${_ORG_ID}','--project-id=${_PROJECT_ID}', '--experiment-id=${_EXPERIMENT_ID}', '--expected-res-score=${_EXPECTED_RES_SCORE}', '--api-key=$$API_KEY' , '--interactive=false', '--monitor=true']

availableSecrets:
  secretManager:
  - versionName: projects/$PROJECT_ID/secrets/x-api-key/versions/latest
    env: API_KEY

substitutions:
  _ACCOUNT_ID: '<ACCOUNT_ID>'
  _ORG_ID: '<ORG_ID>'
  _PROJECT_ID: '<PROJECT_ID>'
  _EXPERIMENT_ID: '<EXPERIMENT_ID>'
  _EXPECTED_RES_SCORE: '100'
```

### Flags

- `--api`: Set the name of the target API (mandatory).
- `--account-id`: Set the account ID (mandatory).
- `--org-id`: Set the organisation id (default "default")
- `--project-id`: Set the HCE project ID (mandatory).
- `--experiment-id`: Set the experiment ID (mandatory for some APIs; a default dummy value is provided).
- `--monitor`: Set this flag true for constantly monitoring the experiment execution with and interval of 2s.
- `--api-key`: Set the API key (mandatory).
- `--interval`: Set the interval provided for multiple iterations (a default value of 2s is provided for some APIs).
- `--timeout`: Set the timeout provided for multiple iterations (a default value of 180s is provided for some APIs).

To know more about Harness Chaos Engineering visit [here](https://developer.harness.io/docs/chaos-engineering).