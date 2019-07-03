# Node.js HTTP Cloud Function

This cloud builder reduces the boilerplate required to build, test and deploy a [Cloud Function](https://cloud.google.com/functions/) that is written using Node.js and intended to respond to an HTTP trigger.

## Cloud Builder Substitutions

| Substitution         | Description                                                                                                                                                                     |                                     |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------- |
| `_NODE_VERSION`      | Accepts a value corresponding to any tag supported by the [official `node` docker image](https://hub.docker.com/_/node/). Will be used as the base image for the cloud builder. | **Default value:** `10-jessie-slim` |
| `_CLOUD_SDK_VERSION` | Version of the Cloud SDK to install when building the cloud builder.                                                                                                            | **Default value:** `251.0.0`        |

## Cloud Builder Parameters

| Parameter                               | Description                                                                     |                                                                                               |
| --------------------------------------- | ------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `--help`                                | List parameter descriptions.                                                    |                                                                                               |
| `--no-build`                            | Do not execute a build.                                                         |                                                                                               |
| `--yarn-build-script=YARN_BUILD_SCRIPT` | The script defined in the `package.json` which builds the project.              | **Default value:** `build`                                                                    |
| `--no-test`                             | Do not execute tests.                                                           |                                                                                               |
| `--yarn-test-script=YARN_TEST_SCRIPT`   | The script defined in the `package.json` which tests the project.               | **Default value:** `test`                                                                     |
| `--deployment-name=DEPLOYMENT_NAME`     | The name that should be used when deploying the Cloud Function.                 | **Required**                                                                                  |
| `--runtime`                             | Runtime in which to run the function.                                           | **Options:** `nodejs6` (deprecated), `nodejs8` or `nodejs10`<br>**Default value:** `nodejs10` |
| `--entry-point=ENTRY_POINT`             | The named export which should be deployed.                                      | **Default value:** The value assigned to `--deployment-name`.                                 |
| `--env-vars-file=ENV_VARS_FILE`         | The file containing environment variables to be supplied to the cloud function. |
