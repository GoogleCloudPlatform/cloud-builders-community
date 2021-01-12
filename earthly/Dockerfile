FROM alpine

RUN apk add --no-cache wget

ARG VERSION=0.4.4
ARG PLATFORM=linux
ARG ARCH=amd64

RUN wget -qO /usr/local/bin/earthly https://github.com/earthly/earthly/releases/download/v$VERSION/earthly-$PLATFORM-$ARCH && \
    chmod +x /usr/local/bin/earthly

ENTRYPOINT ["earthly"]
