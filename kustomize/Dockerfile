FROM gcr.io/cloud-builders/kubectl

RUN wget -qO /builder/kustomize https://github.com/kubernetes-sigs/kustomize/releases/download/v1.0.6/kustomize_1.0.6_linux_amd64

COPY kustomize.bash /builder/kustomize.bash

ENTRYPOINT ["/builder/kustomize.bash"]
