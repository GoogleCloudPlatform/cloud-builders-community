# Cypress dependencies

This allows you to run cypress in GCB.

You need to have `cypress` in you `package.json` and run `yarn install` before running cypress.

## Example build

```yaml
steps:
- name: 'gcr.io/cloud-builders/yarn'
  args: ['install']
- name: 'gcr.io/$PROJECT_ID/cypress-dependencies'
  args: ['cypress', 'run', '--headless']
```

## Useful links
- Cypress documentation: https://docs.cypress.io
- Cypress inside docker: https://docs.cypress.io/examples/examples/docker.html
- Base image: https://hub.docker.com/r/cypress/browsers
