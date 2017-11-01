FROM ubuntu:trusty

ARG version=1.16.1

# https://docs.docker.com/compose/install/
RUN \
   apt-get -y update && \
   apt-get -y install ca-certificates curl && \
   curl -L "https://github.com/docker/compose/releases/download/$version/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
   chmod +x /usr/local/bin/docker-compose

ENTRYPOINT ["/usr/local/bin/docker-compose"]
