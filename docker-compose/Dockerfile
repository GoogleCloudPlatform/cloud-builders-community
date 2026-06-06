FROM ubuntu:focal

ARG DOCKER_COMPOSE_VERSION

# https://docs.docker.com/compose/install/standalone/
# Docker Compose v2.x changed the release binary naming convention:
#   v1.x: docker-compose-Linux-x86_64   (capitalized OS from uname -s)
#   v2.x: docker-compose-linux-x86_64   (lowercase OS)
# Lowercasing the OS name with tr ensures v2.x download URLs resolve correctly.
# --fail causes curl to exit non-zero on HTTP errors (e.g. 404) rather than
# silently writing an HTML error page that would later fail as an invalid binary.
RUN \
   apt-get -y update && \
   apt-get -y install ca-certificates curl docker.io && \
   rm -rf /var/lib/apt/lists/* && \
   OS=$(uname -s | tr '[:upper:]' '[:lower:]') && \
   ARCH=$(uname -m) && \
   curl --fail -L "https://github.com/docker/compose/releases/download/v${DOCKER_COMPOSE_VERSION}/docker-compose-${OS}-${ARCH}" -o /usr/local/bin/docker-compose && \
   chmod +x /usr/local/bin/docker-compose

ENTRYPOINT ["/usr/local/bin/docker-compose"]
