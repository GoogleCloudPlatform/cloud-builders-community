# Harness Chaos

The harness-chaos builder step is used to launch chaos experiments and test the resilience of the applications.

Here’s an example demonstrating how to launch and validate a chaos experiment. You can pass the `--workflow-id` flag along with the `--expected-resilience-score` flag to ensure that the actual resilience score of the experiment run meets the expected threshold, along with other necessary flags.

**Example:**
```
steps:
- name: 'gcr.io/$PROJECT_ID/harness-chaos'
  args: ['generate', '--api=run-and-monitor-experiment', '--account-id=${_ACCOUNT_ID}','--org-id=${_ORG_ID}','--project-id=${_PROJECT_ID}', '--workflow-id=${_EXPERIMENT_ID}', '--expected-resilience-score=${_EXPECTED_RES_SCORE}', '--api-key=${_X_API_KEY}' ]

  env:
    - 'ACCOUNT_ID=${_ACCOUNT_ID}'
    - 'ORG_ID=${_ORG_ID}'
    - 'PROJECT_ID=${_PROJECT_ID}'
    - 'WORKFLOW_ID=${_EXPERIMENT_ID}'
    - 'EXPECTED_RES_SCORE=${_EXPECTED_RES_SCORE}'
    - 'X_API_KEY=${_X_API_KEY}'

substitutions:
  _ACCOUNT_ID: '<ACCOUNT_ID>'
  _ORG_ID: 'ORG_ID'
  _PROJECT_ID: 'PROJECT_ID'
  _EXPERIMENT_ID: 'EXPERIMENT_ID'
  _EXPECTED_RES_SCORE: '100'
  _X_API_KEY: 'X_API_KEY' // required for authorization
```