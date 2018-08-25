# Android Builder
The Dockerfile and scripts here help you use Google Cloud Builder to build Android applications.  
These instructions assume that you have [Android SDK](https://developer.android.com/studio/index.html) installed and can build the subject Android application locally on your workstation.  
Steps to use:

##### 1. Create a list of the android sdk packages you need for your builds and save it to the builder directory as 'packages.txt'.  To generate the list from your workstation a script is provided: 

  ```
  $ ./installed-package-list.sh
  ```
**Note that this script will add every package from the workstation it's run on to the packages.txt file.** As a result, it will probably include things you don't need for your build container.  **Including components you don't need will increase both builder and application build times.**  Therefore, it is recommended that after you generate the packages.txt file you edit it to include only those sdk components that you need. Alternatively you can create this file manually, for example: 

  packages.txt |  
  :-----------|
  extras;m2repository;com;android;support;constraint;constraint-layout;1.0.2   <br/> extras;google;google_play_services <br/> extras;android;gapid;3 <br/>extras;google;auto <br/> extras;android;m2repository <br/>extras;google;webdriver <br/> platform-tools <br/> platforms;android-26 <br/>|


##### 2. Save your android sdk license to a shell variable. 
  ```
  $ ANDROID_SDK_LICENSE=$(tail -1 [ANDROID_SDK_HOME]/licenses/android-sdk-license)
  ```
 
  *Replace `[ANDROID_SDK_HOME]` with the location of your android sdk.*  
 *Usually:* 
  
  OS | Location 
  :---:|:---------:| 
  Mac | ~/Library/android/sdk |
  Unix | ~/Android/Sdk |

  As an alternative you can copy the licenses directory to a _LICENSES_BUCKET
  ```
  $ gsutil rsync -d [ANDROID_SDK_HOME]/licenses/ gs://[_LICENSES_BUCKET]
  ```
  Where _LICENSES_BUCKET is a gs://* URI to a google cloud storage bucket that can be accessed within the project.
  Be sure to use the _LICENSES_BUCKET as a substitution (see below)

##### 3. Submit the build to Google Cloud Build.

  ```
  $ gcloud builds submit --config cloudbuild.yaml . --substitutions=_ANDROID_SDK_LICENSE=$ANDROID_SDK_LICENSE
  ```

  or 
  ```
  gcloud container builds submit --config cloudbuild.yaml . --substitutions=_LICENSES_BUCKET=gs://[_LICENSES_BUCKET]
  ```

##### 4. Copy the example build configuration to the root directory of your android project. 
  
  ```
  $ cp android-build-example/simple-cloudbuild.yaml [ANDROID_PROJECT_DIR]

  ```
##### 5. Create a google cloud storage bucket to cache gradle dependencies. Replace `[ANDROID_BUILD_CACHE]` with the name of the bucket where you want to store the gradle dependencies.

  ```
  $ gsutil mb gs://[ANDROID_BUILD_CACHE]
  ```
##### 6. Create a google cloud storage bucket to deploy your android builds to: 
 
  ```
  $ gsutil mb gs://[DEBUG_BUILD_BUCKET]
  ```

##### 7. Create the build trigger for your android application and repository following [these instructions](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds), *adding the following three Substitution Variables*: 

Variable | Value 
---------------------:|:----------|
`_ANDROID_BUILD_CACHE`  | `[ANDROID_BUILD_CACHE]`
`_ANDROID_SDK_LICENSE`  | `$ANDROID_SDK_LICENSE`
`_LICENSES_BUCKET`      | `[_LICENSES_BUCKET]`
`_DEBUG_BUILD_BUCKET`        | `[DEBUG_BUILD_BUCKET]`
   

#### 8. Get your keystore hashes
If you are using the debug build to run tests that use a third party authentication system you may need the hashes of the debug keystore used to sign the build.  Note that these are deliberately output during the docker build.  To find these values review the output from the submission of the Docker build for this builder. 

# What's included with this builder
## Android Builder Image 
The Dockerfile contained in this repository builds a an Docker container that includes the android SDK and gradle.  Customize this file to change the Android SDK components included in the base image.  

## gradle-build
The gradle-build scripts wraps the gradle build call with bash that will extract a ".gradle" directory that contains cached depenencies and zip up the same directory after the build completes.  This enables the rest of the cloudbuild to cache the gradle dependencies in a Google Cloud Storage bucket. 

## android-build-example/simple-cloudbuild.yaml
This is an example build configuration that uses [Cloud Build](https://cloud.google.com/cloud-build/) to build an Android application debug APK and make it available via a [Cloud Storage](https://cloud.google.com/storage/docs/) bucket.

## android-build-example/fabric-beta-dist-cloudbuild.yaml
This is an example build configuration that uses [Cloud Build](https://cloud.google.com/cloud-build/) to build and distribute an Android application to beta testers using Fabric's Crashlytics. 

## cloudbuild.yaml 
This file builds the android builder.

## ci-cloudbuild.yaml
This file supports testing this cloud builder by adding a step that creates the packages.txt file prior to the build.

## Contact Us

Please use [issue tracker](https://github.com/GoogleCloudPlatform/android-cloud-build/issues)
on GitHub to report any bugs, comments or questions regarding SDK development.

We welcome all usage-related questions on [Stack Overflow](http://stackoverflow.com/questions/tagged/google-android-cloud-build)
tagged with `google-cloud-container-builder`.

## More Information

* [Google Cloud Build](https://cloud.google.com/cloud-build/docs/)

