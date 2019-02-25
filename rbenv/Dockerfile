ARG UBUNTU_VERSION=${UBUNTU_VERSION:-18.04}
ARG RUBY_VERSION=${RUBY_VERSION:-2.6.1}

FROM ubuntu:${UBUNTU_VERSION}

MAINTAINER "Dwi Fahni Denni <zeroc0d3.0912@gmail.com>"

RUN apt-get update; sync
RUN apt-get -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" install \
      build-essential \
      tar \
      bzip2 \
      curl \
      wget \
      libssl-dev \
      libreadline-dev \
      zlib1g-dev \
      gnupg \
      gnupg2 \
      git

COPY rootfs /

RUN chmod +x /tmp/install.sh
RUN /tmp/install.sh

RUN apt-get -y clean

ONBUILD RUN ["ruby", "/tmp/index.rb"]

EXPOSE 22
