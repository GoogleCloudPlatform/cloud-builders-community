# Tool builder: `gcr.io/cloud-builders/rbenv`

This Cloud Build builder runs the `rbenv` (Ruby Environment) tool.

You should consider instead using an [official `ubuntu`
image](https://hub.docker.com/_/ubuntu/) and specifying the `rbenv` entrypoint:

```yaml
steps:
- name: rbenv:18.04-2.6.1
  entrypoint: rbenv
  args: ['install']
```

This allows you to use any supported version of RBENV.

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud builds submit
