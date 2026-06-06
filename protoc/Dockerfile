FROM debian:bookworm-slim

ARG VERS=27.3
ARG ARCH=linux-x86_64

RUN echo "Building protoc Cloud Builder ${VERS}-${ARCH}" && \
    apt-get update -y && apt-get upgrade -y && \
    apt-get install wget unzip -y && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/* && \
    wget --no-verbose "https://github.com/protocolbuffers/protobuf/releases/download/v${VERS}/protoc-${VERS}-${ARCH}.zip" && \
    unzip "protoc-${VERS}-${ARCH}.zip" -d protoc && \
    rm "protoc-${VERS}-${ARCH}.zip"

ENV PATH=$PATH:/protoc/bin/
ENTRYPOINT ["protoc"]
CMD ["--help"]
