FROM launcher.gcr.io/google/ubuntu16_04
ARG DOCTL_VERSION=1.39.0
ARG KUBECTL_VERSION=1.16.6
RUN curl -sL https://github.com/digitalocean/doctl/releases/download/v${DOCTL_VERSION}/doctl-${DOCTL_VERSION}-linux-amd64.tar.gz | tar -xzv -C /usr/local/bin
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl && \ 
    mv ./kubectl /usr/local/bin
COPY kubectl.bash /builder/kubectl.bash
RUN chmod +x /builder/kubectl.bash
ENTRYPOINT ["/builder/kubectl.bash"]
