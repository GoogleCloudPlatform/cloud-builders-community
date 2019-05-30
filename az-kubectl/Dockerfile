FROM gcr.io/learning-resources-hub/az

# Install kubectl component
RUN /usr/local/bin/az aks install-cli

COPY kubectl.bash /builder/kubectl.bash

ENTRYPOINT ["/builder/kubectl.bash"]

