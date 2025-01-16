ARG PY_VERSION=3.11
FROM python:${PY_VERSION}-bookworm

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get -y upgrade \
    && apt-get -y install --no-install-recommends  \
    curl ca-certificates

ENV PATH="/root/.local/bin:$PATH"
RUN curl --proto '=https' --tlsv1.2 -fsSL https://static.pantsbuild.org/setup/get-pants.sh | bash

ENTRYPOINT [ "pants" ]
