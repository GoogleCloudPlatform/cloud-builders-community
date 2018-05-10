FROM gcr.io/cloud-builders/gsutil

ENV TERRAFORM_VERSION=0.11.7
ENV TERRAFORM_VERSION_SHA256SUM=6b8ce67647a59b2a3f70199c304abca0ddec0e49fd060944c26f666298e23418

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