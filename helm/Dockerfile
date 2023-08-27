FROM gcr.io/google.com/cloudsdktool/cloud-sdk:slim

ARG HELM_VERSION=v3.12.0
ENV HELM_VERSION=$HELM_VERSION

COPY kubectl.bash /builder/kubectl.bash
COPY helm.bash /builder/helm.bash

RUN chmod +x /builder/helm.bash /builder/kubectl.bash && \
  mkdir -p /builder/helm && \
  apt-get update && \
  apt-get install -y --no-install-recommends curl kubectl google-cloud-sdk-gke-gcloud-auth-plugin && \
  curl -SL https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz -o helm.tar.gz && \
  tar zxvf helm.tar.gz --strip-components=1 -C /builder/helm linux-amd64 && \
  rm helm.tar.gz && \
  apt-get --purge -y autoremove && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

ENV PATH=/builder/helm/:$PATH
ENV USE_GKE_GCLOUD_AUTH_PLUGIN=True

ENTRYPOINT ["/builder/helm.bash"]
