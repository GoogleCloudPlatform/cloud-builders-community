FROM ubuntu

ARG PACKER_VERSION=1.2.4
ARG PACKER_VERSION_SHA256SUM=258d1baa23498932baede9b40f2eca4ac363b86b32487b36f48f5102630e9fbb

RUN apt-get -y update && apt-get -y install ca-certificates && rm -rf /var/lib/apt/lists/*

COPY packer_${PACKER_VERSION}_linux_amd64.zip .
RUN echo "${PACKER_VERSION_SHA256SUM} packer_${PACKER_VERSION}_linux_amd64.zip" > checksum && sha256sum -c checksum
RUN apt-get -y update && apt-get -y install unzip && /usr/bin/unzip packer_${PACKER_VERSION}_linux_amd64.zip

FROM ubuntu
RUN apt-get -y update && apt-get -y install ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=0 packer /usr/bin/packer
ENTRYPOINT ["/usr/bin/packer"]
