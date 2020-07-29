FROM alpine:3.12

ARG VERS=3.12.4
ARG ARCH=linux-x86_64

RUN echo "Building protoc Cloud Builder ${VERS}-${ARCH}" && \
    apk --update add --no-cache wget unzip && \
    rm -rf /var/cache/apk/* && \
    wget "https://github.com/protocolbuffers/protobuf/releases/download/v${VERS}/protoc-${VERS}-${ARCH}.zip" && \
    unzip "protoc-${VERS}-${ARCH}.zip" -d protoc && \
    rm "protoc-${VERS}-${ARCH}.zip"

ENV PATH=$PATH:/protoc/bin/
ENTRYPOINT ["protoc"]
CMD ["--help"]
