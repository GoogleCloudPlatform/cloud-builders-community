ARG BASE_IMAGE=gcr.io/cloud-builders/java/javac:8
FROM ${BASE_IMAGE}

ARG SBT_VERSION=1.3.12
ARG SHA=1914caf18790d2320c37ef091ba696f8fc00808e034597e948196d0ec0f39e0a
ARG BASE_URL=https://github.com/sbt/sbt/releases/download

RUN apt-get update -qqy \
  && apt-get install -qqy curl bc \
  && mkdir -p /usr/share \
  && curl -fsSL -o "sbt-${SBT_VERSION}.zip" "${BASE_URL}/v${SBT_VERSION}/sbt-${SBT_VERSION}.zip" \
  && echo "${SHA}  sbt-${SBT_VERSION}.zip" | sha256sum -c - \
  && unzip -qq "sbt-${SBT_VERSION}.zip" \
  && rm -f "sbt-${SBT_VERSION}.zip" \
  && mv sbt "/usr/share/sbt-${SBT_VERSION}" \
  && ln -s "/usr/share/sbt-${SBT_VERSION}/bin/sbt" /usr/bin/sbt \
  && apt-get remove -qqy --purge curl \
  && rm /var/lib/apt/lists/*_*

ENTRYPOINT ["/usr/bin/sbt"]
