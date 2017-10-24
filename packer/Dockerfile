FROM ubuntu
RUN apt-get -y update && apt-get -y install ca-certificates && rm -rf /var/lib/apt/lists/*
COPY packer_1.0.0_linux_amd64.zip .
RUN echo 'ed697ace39f8bb7bf6ccd78e21b2075f53c0f23cdfb5276c380a053a7b906853  packer_1.0.0_linux_amd64.zip' > checksum && sha256sum -c checksum
RUN apt-get -y update && apt-get -y install unzip && /usr/bin/unzip packer_1.0.0_linux_amd64.zip

FROM ubuntu
RUN apt-get -y update && apt-get -y install ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=0 packer /usr/bin/packer
ENTRYPOINT ["/usr/bin/packer"]
