FROM alpine:latest

LABEL maintainer "Ram Gopinathan"

ARG SONARQUBE_SCANNER_CLI_VERSION="3.2.0.1227"

ENV SONARQUBE_SCANNER_HOME /opt/sonar-scanner-${SONARQUBE_SCANNER_CLI_VERSION}-linux
ENV SONARQUBE_SCANNER_BIN ${SONARQUBE_SCANNER_HOME}/bin
ENV SONAR_SCANNER_CLI_DOWNLOAD_URL "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-${SONARQUBE_SCANNER_CLI_VERSION}-linux.zip"

RUN apk update \
	&& apk upgrade \
	&& apk add ca-certificates \
	&& update-ca-certificates \
	&& apk add --update openjdk8-jre tzdata curl unzip bash \
	&& rm -rf /var/cache/apk/* \
    && mkdir -p /tmp/sonar-scanner  \
	&& curl -L --silent ${SONAR_SCANNER_CLI_DOWNLOAD_URL} >  /tmp/sonar-scanner/sonar-scanner-cli-${SONARQUBE_SCANNER_CLI_VERSION}-linux.zip  \
    && mkdir -p /opt  \
	&& unzip /tmp/sonar-scanner/sonar-scanner-cli-${SONARQUBE_SCANNER_CLI_VERSION}-linux.zip -d /opt  \
	&& rm -rf /tmp/sonar-scanner


ENV PATH $PATH:$SONARQUBE_SCANNER_BIN

COPY launch.sh /

WORKDIR ${SONARQUBE_SCANNER_HOME}

ENTRYPOINT ["/launch.sh"]