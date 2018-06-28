# [Flutter](https://flutter.io) Cloud Builder Build Step
========================================================

This container builder provides the `[flutter]`(https://github.com/flutter/flutter) tool.

## Building the Builder

To build this builder, run the following command in this directory.

```
gcloud container builds submit . --config=cloudbuild.yaml
```

## Using the Flutter Build Step

* To create a release APK of a Flutter application on Android.

```
- name: 'gcr.io/$PROJECT_ID/flutter'
  args: [ 'build', 'apk', '--release' ]
```

* To run all Flutter unit tests.

```
 - name: 'gcr.io/$PROJECT_ID/flutter'
   args: [ 'test' ]
```