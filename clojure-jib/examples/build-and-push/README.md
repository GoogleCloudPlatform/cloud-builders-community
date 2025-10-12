# Clojure Jib example

This directory contains a simple clojure project that uses jetty to start up a web server.

This assumes you have already built the `clojure-jib` build step and pushed it to
`gcr.io/$PROJECT_ID/clojure-jib`.

## Executing the builder

Inside this directory, run the following command:

```
gcloud builds submit --config=cloudbuild.yaml .
```

Build output will be displayed.  At the end you should see the following:

```
Registry image: gcr.io/personalsdm-216019/clojure-jib-example
jib:WARN      Base image 'gcr.io/distroless/java' does not use a specific image digest - build may not be reproducible
jib:PROGRESS  Getting manifest for base image gcr.io/distroless/java...
jib:PROGRESS  Building dependencies layer layer...
jib:PROGRESS  Building clojure application layer layer...
jib:LIFECYCLE Using base image with digest: sha256:629d4fdc17eec821242d45497abcb88cc0442c47fd5748baa79d88dde7da3e2d
jib:LIFECYCLE
jib:LIFECYCLE Container entrypoint set to [java, -Dclojure.main.report=stderr, -Dfile.encoding=UTF-8, -cp, src:lib/clojure -1.10.3.jar:lib/core.specs.alpha-0.2.56.jar:lib/spec.alpha-0.2.194.jar, clojure.main, -m, example.main]
PUSH
DONE
--------------------------------------------------------------------------------------------------------------------------
ID                                    CREATE_TIME                DURATION  SOURCE                                          IMAGES  STATUS
08db0c7b-451e-4358-b094-8c14a1e93ff7  2022-04-04T04:55:54+00:00  33S       gs://personalsdm-216019_cloudbuild/source/16490 48153.567737-792fa775422f45e08d963a6b504555d5.tgz  -       SUCCESS
```

## CloudRun your new image

In the previous step, you built and pushed a tiny clojure application to your project's google container registry.  You can deploy this to CloudRun and try it out using the following.

```bash
gcloud run deploy clojure-jib-example --image gcr.io/$PROJECT_ID/clojure-jib-example --allow-unauthenticated --region us-east1
```

The output will give you a url that you can curl to see your application run.  The command above requires that you set the `PROJECT_ID` environment variable and will deploy in the us-east1 region.

```
Deploying container to Cloud Run service [clojure-jib-example] in project [personalsdm-216019] region [us-east1]
✓ Deploying... Done.
  ✓ Creating Revision...
  ✓ Routing traffic...
  ✓ Setting IAM Policy...
Done.
Service [clojure-jib-example] revision [clojure-jib-example-00002-bir] has been deployed and is serving 100 percent of traffic.
Service URL: https://clojure-jib-example-bq55q5z23q-ue.a.run.app
```

Try curling your `Service URL`.  The first invocation will have take extra time as your CloudRun instance warms up.

Remember to delete your service afterwards.

```bash
gcloud run services delete clojure-jib-example --region us-east1
```
