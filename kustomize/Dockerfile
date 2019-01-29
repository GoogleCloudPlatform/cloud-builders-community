FROM gcr.io/cloud-builders/gcloud

ENV VER 1.0.11
ENV VERSION v${VER}

COPY kustomize.bash /builder/kustomize.bash

RUN apt-get update && \
    apt-get install -y wget && \
    wget https://github.com/kubernetes-sigs/kustomize/releases/download/${VERSION}/kustomize_${VER}_linux_amd64 && \
    mkdir /builder/kustomize && \
    mv kustomize_${VER}_linux_amd64 /builder/kustomize/kustomize && \
    chmod +x /builder/kustomize/kustomize && \
    chmod +x /builder/kustomize.bash && \
    apt-get remove --purge -y wget && \
    apt-get --purge -y autoremove && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV PATH=/builder/kustomize:$PATH

ENTRYPOINT ["/builder/kustomize.bash"]
