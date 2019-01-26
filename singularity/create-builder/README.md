# Step 1. Create a Singularity builder

This is the first step that you will need to do to use [Google Cloud Build](https://cloud.google.com/cloud-build/docs/)
to build Singularity containers. Here we will define a 
[custom build step](https://cloud.google.com/cloud-build/docs/create-custom-build-steps)
to create a builder. The files needed to do this are:

 * [builder.yaml](builder.yaml): this is your Cloud Builder recipe! It's a configuration file you will submit to build the custom build step
 * [Dockerfile](Dockerfile): is the container that we will install Singularity in, and then eventually use when we [run our builder](../run-builder) to build the Singularity container. 

## What Singularity Version?

The container provided here will support Singularity builds for version 3.0.0 and up. 
As of Singularity 3.0.0, the software is implemented in GoLang and follows an installation
procedure appropriate for that. We recommend using these latest versions of Singularity.

### Build your Builder

We assume that you have installed gcloud and have your project set up, and APIs enabled.
You can create your builder by targeting the builder.yaml configuration, and specifying 
your Singularity version. Here we supply version 3.0.2, the most recent stable release
at the time of this writing.

```bash
$ gcloud builds submit --config=builder.yaml --substitutions=_SINGULARITY_VERSION="3.0.2" .
```

This verison should coincide a [release tag](https://github.com/sylabs/singularity/releases) on 
the [sylabs/singularity](https://github.com/sylabs/singularity) repository.

Once you run the above, you should see output in your terminal for building the container,
and success that it's complete! Now you can move on to [using the builder](../run-builder).
