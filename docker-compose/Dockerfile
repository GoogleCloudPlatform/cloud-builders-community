FROM ubuntu:bionic

ARG version=1.24.0

# https://docs.docker.com/compose/install/
RUN \
   apt -y update && \
   apt -y install ca-certificates curl docker.io && \
   curl -L "https://github.com/docker/compose/releases/download/$version/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
   chmod +x /usr/local/bin/docker-compose

ENTRYPOINT ["/usr/local/bin/docker-compose"]
