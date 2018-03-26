FROM gcr.io/cloud-builders/kubectl

RUN mkdir -p /builder/bin && \
	curl -sSLo /builder/bin/skaffold https://storage.googleapis.com/skaffold/v0.1.0/skaffold-linux-amd64 && \
	chmod +x /builder/bin/skaffold
ENV PATH=/builder/bin/:$PATH

COPY skaffold.bash /builder/skaffold.bash

ENTRYPOINT ["/builder/skaffold.bash"]
