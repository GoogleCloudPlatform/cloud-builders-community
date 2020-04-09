# Python Hello World

This demonstrates a simple project for printing a "Hello World!" with a
Python runtime environment provided through the [ActiveState Platform].

## Usage

Before you can run this example you'll need to set up a Python3 project
on the [ActiveState Platform], to do this reference the [Creating Custom Projects]
documentation. Ensure that you create project that uses a Linux 64bit platform.

Once you have created your project update the Dockerfile and the
activestate.yaml to use the namespace for your project rather than the
placeholder value.

This example uses secrets, read the *Private Projects and State Secrets*
in the [main readme] on how to set up cloud build to use State Secrets.

## Executing Your New Builder

Once configured you should be able to simply run

```
gcloud builds submit --config=cloudbuild.yaml
```

   [ActiveState Platform]: https://www.activestate.com/products/platform/
   [Creating Custom Projects]: https://docs.activestate.com/platform/projects/custom-builds/
   [main readme]: ../../README.md