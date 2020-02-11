FROM gcr.io/cloud-builders/gcloud

ARG HELM_VERSION=v2.16.1
ENV HELM_VERSION=$HELM_VERSION

COPY helm.bash /builder/helm.bash

RUN chmod +x /builder/helm.bash && \
  mkdir -p /builder/helm && \
  apt-get update && \
  apt-get install -y curl && \
  curl -SL https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz -o helm.tar.gz && \
  tar zxvf helm.tar.gz --strip-components=1 -C /builder/helm linux-amd64 && \
  rm helm.tar.gz && \
  apt-get --purge -y autoremove && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

ENV PATH=/builder/helm/:$PATH

ENTRYPOINT ["/builder/helm.bash"]
