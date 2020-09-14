FROM gcr.io/cloud-builders/gcloud

# overridden in cloudbuild.yaml
ARG VER=v3.8.1
ENV VERSION ${VER}

COPY kustomize.bash /builder/kustomize.bash

RUN apt-get update && \
  apt-get install -y wget && \
  wget -nv https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F${VERSION}/kustomize_${VERSION}_linux_amd64.tar.gz && \
  tar xvzf kustomize_${VERSION}_linux_amd64.tar.gz && \
  mkdir /builder/kustomize && \
  mv kustomize /builder/kustomize/kustomize && \
  chmod +x /builder/kustomize/kustomize && \
  chmod +x /builder/kustomize.bash && \
  rm kustomize_${VERSION}_linux_amd64.tar.gz && \
  apt-get remove --purge -y wget && \
  apt-get --purge -y autoremove && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

ENV PATH=/builder/kustomize:$PATH

ENTRYPOINT ["/builder/kustomize.bash"]
