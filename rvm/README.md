# Tool builder: `gcr.io/cloud-builders/rvm`

This Cloud Build builder runs the `rvm` (Ruby Version Manager) tool.

You should consider instead using an [official `ubuntu`
image](https://hub.docker.com/_/ubuntu/) and specifying the `rvm` entrypoint:

```yaml
steps:
- name: rvm:18.04-2.6.1
  entrypoint: rvm
  args: ['install']
```

This allows you to use any supported version of RBENV.

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud builds submit
