# protoc

This tool defines a custom build step that allows the Cloud Build worker to
run the
[protocol buffer compiler](https://github.com/protocolbuffers/protobuf), `protoc`.

## When to use this builder

The gcr.io/cloud-builders/protoc build step should be used when you want to run
`protoc` as part of your (Google Cloud) build process.

## Building this builder

To build this builder and push the resulting image to the Container Registry *in
your project*, run the following command in this directory:

    $ gcloud builds submit . --config=cloudbuild.yaml

## About the Dockerfile

The Docker file is where the `protoc` command is installed, and is the more complex of the two files: 

    FROM ubuntu
    
    ARG PROTOC_VERSION=3.6.1
    ARG PROTOC_TARGET=linux-x86_64
    ARG ASSET_NAME=protoc-${PROTOC_VERSION}-${PROTOC_TARGET}.zip
    
    RUN apt-get -qy update && apt-get -qy install python wget unzip && rm -rf /var/lib/apt/lists/*
    
    RUN echo "${PROTOC_VERSION}/${ASSET_NAME}"
    
    RUN wget https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${PROTOC_TARGET}.zip && \
    unzip ${ASSET_NAME} -d protoc && rm ${ASSET_NAME}
    
    ENV PATH=$PATH:/protoc/bin/
    ENTRYPOINT ["protoc"]
    CMD ["--help]

Breaking this down:

 - Define the first read only layer of the final "protoc" image we want to
end-up with. I chose Ubuntu because it's the what I run locally. Any minimal
Linux "base image" will do but it must have the following binaries installed:
`apt-get`, `wget`, `unzip`, and `rm`:

`FROM ubuntu`
 
 - Set-up some variables that users can pass at build-time to the builder with
the docker build command using the `--build-arg <varname>=<value>` flag:

`ARG PROTOC_VERSION=3.6.1`
`ARG PROTOC_TARGET=linux-x86_64`
`ARG ASSET_NAME=protoc-${PROTOC_VERSION}-${PROTOC_TARGET}.zip`

 - Run `apt-get -qy update` to "resynchronize the package index files from their
sources". `q` omits progress indicators, `y` assumes yes as an answer to any
prompts encountered:

`RUN apt-get -qy update`

 - Install **Python**, **Wget** (retrieves content from web server),
and **Unzip**.

`RUN apt-get -qy install python wget unzip`

 - Remove any files created as part of the previous steps (and that are no
longer needed):

`RUN rm -rf /var/lib/apt/lists/*`

The previous three [RUN](https://docs.docker.com/engine/reference/builder/#run) instructions can be combined into one:

`RUN apt-get -qy update && apt-get -qy install python wget unzip && rm -rf /var/lib/apt/lists/*`

 - Use the [ENV](https://docs.docker.com/engine/reference/builder/#env) instruction to update the `PATH` environment
 to include the location of the `protoc` binary in the final environment (image). 

`ENV PATH=$PATH:/protoc/bin/`

Set the [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) of the image such that the image runs
as a `protoc` executable. Not, since the previous step added `protoc` to `$PATH`, we need only specify the binary to run
(not the full path):

`ENTRYPOINT ["protoc"]`

 - Use the [CMD](https://docs.docker.com/engine/reference/builder/#cmd) instruction so that if no options are provided
when running the `protoc` image, `protoc --help` will run:

`CMD ["--help]`
