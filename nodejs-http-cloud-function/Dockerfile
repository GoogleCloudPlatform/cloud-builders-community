ARG NODE_VERSION

FROM docker:17.12.0-ce as static-docker-source
FROM node:$NODE_VERSION

COPY --from=static-docker-source /usr/local/bin/docker /usr/local/bin/docker

ARG CLOUD_SDK_VERSION
ENV CLOUD_SDK_VERSION=$CLOUD_SDK_VERSION

ENV PATH "$PATH:/opt/google-cloud-sdk/bin/"

RUN apt-get -qqy update && apt-get install -qqy \
    curl \
    gcc \
    python-dev \
    python-setuptools \
    apt-transport-https \
    lsb-release \
    openssh-client \
    git \
    gnupg && \
    easy_install -U pip && \
    pip install -U crcmod && \
    export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb https://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" > /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update && \
    apt-get install -y google-cloud-sdk=${CLOUD_SDK_VERSION}-0 && \
    gcloud --version && \
    docker --version

VOLUME ["/root/.config"]

COPY build-test-and-deploy.sh /bin
ENTRYPOINT ["/bin/build-test-and-deploy.sh"]