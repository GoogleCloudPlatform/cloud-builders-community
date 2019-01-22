FROM ubuntu

ARG RKE_VERSION=0.1.15
ARG RKE_VERSION_SHA256SUM=6687ea0d683a29d6c569cf6ffc7b3a522425b83a1a0b7ea2d402befbd495edaa

RUN apt-get -y update && apt-get -y install ca-certificates && rm -rf /var/lib/apt/lists/*

COPY rke_linux-amd64 /usr/bin/rke
RUN echo "${RKE_VERSION_SHA256SUM} /usr/bin/rke" > checksum && sha256sum -c checksum
RUN chmod +x /usr/bin/rke

ENTRYPOINT ["/usr/bin/rke"]
