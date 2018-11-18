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
