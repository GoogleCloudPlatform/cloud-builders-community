FROM gcr.io/cloud-builders/gcloud

RUN apt-get update && apt-get -y install make \
    && pip install setuptools \
    && git clone https://github.com/GoogleCloudPlatform/deploymentmanager-samples \
    && cd deploymentmanager-samples/ \
    && cd community/cloud-foundation \
    && make prerequisites \
    && make build \
    && make install \
    && cd / \
    && rm -rf /deploymentmanager-samples

ENTRYPOINT ["/usr/local/bin/cft"]
