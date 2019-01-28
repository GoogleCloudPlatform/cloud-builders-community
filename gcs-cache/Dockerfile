FROM gcr.io/cloud-builders/gcloud-slim

RUN apt-get -y update && \
    apt-get -y install zip && \
    apt-get -y install unzip

COPY gcs-cache.bash /builder/gcs-cache.bash

RUN chmod +x /builder/gcs-cache.bash

ENTRYPOINT ["/builder/gcs-cache.bash"]
