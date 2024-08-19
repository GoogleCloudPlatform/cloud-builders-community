FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:alpine

ARG HELM_VERSION=v3.15.4
ENV HELM_VERSION=$HELM_VERSION
ENV USE_GKE_GCLOUD_AUTH_PLUGIN=True

COPY helm.bash /builder/helm.bash

RUN chmod +x /builder/helm.bash && \
  mkdir -p /builder/helm && \
  curl -SL https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz | tar zxv --strip-components=1 -C /builder/helm linux-amd64 && \
  gcloud -q components install gke-gcloud-auth-plugin

ENV PATH=/builder/helm/:$PATH

ENTRYPOINT ["/builder/helm.bash"]
