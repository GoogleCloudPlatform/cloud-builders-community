FROM gcr.io/cloud-builders/git

RUN apt-get update && \
  apt-get install -y wget && \
  wget https://github.com/github/hub/releases/download/v2.3.0-pre10/hub-linux-amd64-2.3.0-pre10.tgz && \
  tar xvfz hub-linux-amd64-2.3.0-pre10.tgz && \
  rm hub-linux-amd64-2.3.0-pre10.tgz && \
  mv hub-linux-amd64-2.3.0-pre10/bin/hub /usr/bin/ && \
  chmod +x /usr/bin/hub && \
  alias git=hub

ENTRYPOINT ["/usr/bin/hub"]
