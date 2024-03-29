FROM gcr.io/cloud-builders/gcloud

ARG HELM_VERSION=v3.14.0
ARG HELMFILE_VERSION=v0.161.0

COPY helmfile.bash /builder/helmfile.bash

# install curl
RUN apt-get update && \
  apt-get install -y curl && \
  apt-get --purge -y autoremove && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

# install helm, helmfile and helm-diff plugin
RUN chmod +x /builder/helmfile.bash && \
  mkdir -p /builder/helmfile && \
  curl -SL https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz -o helm.tar.gz && \
  tar zxvf helm.tar.gz --strip-components=1 -C /builder/helmfile linux-amd64 && \
  rm helm.tar.gz && \
  curl -fSsL https://github.com/helmfile/helmfile/releases/download/${HELMFILE_VERSION}/helmfile_${HELMFILE_VERSION#v}_linux_amd64.tar.gz > /tmp/helmfile.tar.gz && \
  tar zxvf /tmp/helmfile.tar.gz -C /builder/helmfile helmfile && \
  rm /tmp/helmfile.tar.gz && \
  chmod 700 /builder/helmfile/helmfile && \
  /builder/helmfile/helm plugin install https://github.com/databus23/helm-diff

ENV PATH=/builder/helmfile/:$PATH

ENTRYPOINT ["/builder/helmfile.bash"]
