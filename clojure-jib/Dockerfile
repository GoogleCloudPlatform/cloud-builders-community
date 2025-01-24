FROM gcr.io/cloud-builders/gcloud

LABEL maintainer="Jim Clark <jim@atomist.com>"
ARG CLOJURE_VERSION="1.10.3.1029"

RUN curl -O https://download.clojure.org/install/linux-install-$CLOJURE_VERSION.sh && \
    sha256sum linux-install-$CLOJURE_VERSION.sh && \
    echo "fe495d448a8c4565bb37f04481f6231575114a3d918cadb93be6770f43863ad4 *linux-install-$CLOJURE_VERSION.sh" | sha256sum -c - && \
    chmod +x linux-install-$CLOJURE_VERSION.sh && \
    ./linux-install-$CLOJURE_VERSION.sh && \
    clojure -e "(clojure-version)" && \
    clojure -Ttools install io.github.atomisthq/jibbit '{:git/sha "ccc2131aabc81f7391d567772020a6baa729f953"}' :as jib

ENTRYPOINT ["clojure", "-Tjib"]
