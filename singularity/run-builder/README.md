# Step 2: Build a Container

In the next steps, we will use the builder generated in [the first step](../create-builder)
to build a custom container and save to a Google Cloud Bucket. If you have a bucket in
mind, great! If not, then you should create one first:

## Create a Bucket

We can use gsutil to quickly make a bucket. Let's say that our project is called
`container-pancakes`, we would first export it to an environment variable `$PROJECT`
and then create the bucket using the variable:

```bash
$ PROJECT=container-pancakes
$ gsutil mb gs://${PROJECT}-singularity
```

## Build Away!

Now we are ready to build! If you've never used Singularity before, it uses a [Singularity Definition File](https://github.com/sylabs/singularity-userdocs/blob/master/definition_files.rst) (akin to a Dockerfile) as a build definition or recipe.
The sections in the file determine how the container is built. For example:

 - %post: is a chunk of commands that are run after the base operating system is added
 * %environment: is where you can define and export environment variables. It actually turns into a script that is sourced.
 * %runscript: is a snippet of code that is executed when the container is run. If you are familiar with Docker, this is similar to an ENTRYPOINT / CMD.

To give you a dummy recipe, we've provided the [Singularity](Singularity) file in this
folder. The file will use a Docker container as a base (`Bootstrap: docker` and `From: ubuntu:16.04`)
and then install a program that will print fun text to the screen. We will interact with the Google Cloud Builder 
in almost the same way, but provide the [cloudbuild.yaml](cloudbuild.yaml) file instead of the previous
builder.yaml. The resulting container will be saved to Google Cloud Storage. We only need to 
provide the Singularity Version to reference the builder that we [previously built](../create-builder).

### Customize the Configuration File

Take a look quickly at [cloudbuild.yaml](cloudbuild.yaml). 

```yaml
steps:
- name: gcr.io/$PROJECT_ID/singularity-${_SINGULARITY_VERSION}
  args: ['build', 'greeting.sif', 'Singularity']
artifacts:
  objects:
    location: 'gs://${PROJECT_ID}-singularity'
    paths: ['greeting.sif']
```

Notice that there are a few ways you can quickly customize it!

 - *args* contains the name of our Singularity recipe (Singularity) and the container we want to build. Feel free to change these! You can also take advantage of other Singularity build arguments.
 - *artifacts* references our Google Cloud Storage bucket, and the path we want to save to. We have a simple example here, but you could also derive a path that includes versioning (via environment variables, or similar).

### Run the Build

When you are happy with the configuration, build away!

```bash
$ gcloud builds submit --config=cloudbuild.yaml --substitutions=_SINGULARITY_VERSION="3.0.2" .
```

Once the build completes verify that the container with created using the gsutil list (ls) command:

```bash
$ gsutil ls gs://${PROJECT}-singularity/greeting.sif
```

You can then do the equivalent of pulling the container by copying from the bucket to your
local machine:


```bash
$ gsutil cp gs://${PROJECT}-singularity/greeting.sif .
```

### Run the Container

Here is how you can run the container:

```bash
$ singularity run greeting.sif
                                      
__   ____ _ _ __   ___  ___ ___  __ _ 
\ \ / / _` | '_ \ / _ \/ __/ __|/ _` |
 \ V / (_| | | | |  __/\__ \__ \ (_| |
  \_/ \__,_|_| |_|\___||___/___/\__,_|

```

Without any arguments, it prints your username! If you provide command line arguments,
it will render those instead.

```bash
$ singularity run greeting.sif hello tacos!
 _          _ _         _                      _ 
| |__   ___| | | ___   | |_ __ _  ___ ___  ___| |
| '_ \ / _ \ | |/ _ \  | __/ _` |/ __/ _ \/ __| |
| | | |  __/ | | (_) | | || (_| | (_| (_) \__ \_|
|_| |_|\___|_|_|\___/   \__\__,_|\___\___/|___(_)
                                                 
```


### Build Locally

If you are familiar with Singularity and have it installed on your local machine,
you can also build the container locally:

```bash
$ sudo singularity build greeting.sif Singularity
```

If you have any questions, please [open an issue](https://github.com/GoogleCloudPlatform/cloud-builders-community/issues).
