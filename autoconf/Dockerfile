FROM debian

RUN apt-get update && apt-get install -y \
  autoconf \
  build-essential \
  ca-certificates \
  curl \
  git \
  libtool \
  pkg-config \
  wget \
  ;

ENTRYPOINT ["autoconf"]
