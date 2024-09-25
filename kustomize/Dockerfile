FROM alpine:3.12 AS build

ARG KUSTOMIZE_VERSION

RUN apk add --no-cache wget tar
RUN wget -nv https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F${KUSTOMIZE_VERSION}/kustomize_${KUSTOMIZE_VERSION}_linux_amd64.tar.gz && \
  tar xvzf kustomize_${KUSTOMIZE_VERSION}_linux_amd64.tar.gz


FROM gcr.io/cloud-builders/gcloud

ENV PATH=/builder/bin/:$PATH
COPY --from=build kustomize /builder/bin/kustomize
COPY kustomize.bash /builder/kustomize.bash
ENTRYPOINT ["/builder/kustomize.bash"]
