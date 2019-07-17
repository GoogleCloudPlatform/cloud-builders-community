FROM ubuntu:19.10

ARG VERS=3.8.0
ARG ARCH=linux-x86_64

RUN apt-get -qy update && \
    apt-get -qy install wget unzip && \
    rm -rf /var/lib/apt/lists/*

RUN echo "${VERS}-${ARCH}"

RUN wget https://github.com/google/protobuf/releases/download/v${VERS}/protoc-${VERS}-${ARCH}.zip && \
    unzip protoc-${VERS}-${ARCH}.zip -d protoc && \
    rm protoc-${VERS}-${ARCH}.zip

ENV PATH=$PATH:/protoc/bin/
ENTRYPOINT ["protoc"]
CMD ["--help"]
