# Clojure Jib Builder

This builder is used to package a deps.edn clojure project into a container image using [Jib](https://github.com/GoogleContainerTools/jib).  It uses the the [jibbit tool](https://github.com/atomisthq/jibbit) to translate a deps.edn file into a jib build plan.

## Building this builder

To build this, run the following command in this directory.

```bash
gcloud builds submit . --config=cloudbuild.yaml
```

This will push the clojure-jib builder to your project's registry.

## Examples

You can use this builder to package your clojure deps.edn project into a container image and push it to gcr or gar.  There is a sample clojure project in the `examples/build-and-push` directory.  After completing the `Building this builder` step, you can run the build in the `examples/build-and-push` directory and try running the sample clojure project in CloudRun.
