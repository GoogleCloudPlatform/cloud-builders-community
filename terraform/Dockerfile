FROM gcr.io/cloud-builders/gcloud

ENV TERRAFORM_VERSION=0.11.10
ENV TERRAFORM_VERSION_SHA256SUM=43543a0e56e31b0952ea3623521917e060f2718ab06fe2b2d506cfaa14d54527

RUN apt-get update && \
    apt-get -y install curl unzip ca-certificates && \
    curl https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
      > terraform_linux_amd64.zip && \
    echo "${TERRAFORM_VERSION_SHA256SUM} terraform_linux_amd64.zip" > terraform_SHA256SUMS && \
    sha256sum -c terraform_SHA256SUMS --status && \
    unzip terraform_linux_amd64.zip -d /builder/terraform && \
    rm -f terraform_linux_amd64.zip

ENV PATH=/builder/terraform/:$PATH
COPY entrypoint.bash /builder/entrypoint.bash
ENTRYPOINT ["/builder/entrypoint.bash"]