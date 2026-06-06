FROM ubuntu:22.04

RUN \
    apt-get update -y && \
    apt-get -y install --no-install-recommends build-essential git ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    git clone https://github.com/google/jsonnet.git && \
    cd jsonnet && \
    make && \
    ./jsonnet --help

ENTRYPOINT ["/jsonnet/jsonnet"]
