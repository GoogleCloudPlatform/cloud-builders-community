FROM alpine:3.9

ENV TERRAFORM_VERSION=0.12.6
ENV TERRAFORM_VERSION_SHA256SUM=6544eb55b3e916affeea0a46fe785329c36de1ba1bdb51ca5239d3567101876f

COPY terraform_${TERRAFORM_VERSION}_linux_amd64.zip terraform_${TERRAFORM_VERSION}_linux_amd64.zip
RUN echo "${TERRAFORM_VERSION_SHA256SUM}  terraform_${TERRAFORM_VERSION}_linux_amd64.zip" > checksum && sha256sum -c checksum
RUN unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip


FROM gcr.io/cloud-builders/gcloud

ENV PATH=/builder/terraform/:$PATH
COPY --from=0 terraform /usr/bin/terraform
COPY entrypoint.bash /builder/entrypoint.bash
ENTRYPOINT ["/builder/entrypoint.bash"]
