ARG ALPINE_VERSION
ARG NODE_VERSION
FROM node:$NODE_VERSION-alpine$ALPINE_VERSION
ARG ALPINE_VERSION
ARG CHROME_VERSION

# Installs latest Chromium available in the alpine $ALPINE_VERSION branch.
RUN apk update && apk upgrade && \
    echo @$ALPINE_VERSION http://nl.alpinelinux.org/alpine/$ALPINE_VERSION/community >> /etc/apk/repositories && \
    echo @$ALPINE_VERSION http://nl.alpinelinux.org/alpine/$ALPINE_VERSION/main >> /etc/apk/repositories && \
    apk add --no-cache \
      chromium \
      nss \
      freetype \
      freetype-dev \
      harfbuzz \
      ca-certificates \
      ttf-freefont \
      git

# Tell Puppeteer to skip installing Chrome. We'll be using the installed package.
ENV PUPPETEER_SKIP_CHROMIUM_DOWNLOAD true
ENV JEST_PUPPETEER_CONFIG jest-puppeteer.config.ci.js
ENV CI true

# Install puppeteer version compatible with the Chromium installed above.
# - This installation of puppeteer will not be used if the web project this cloud builder is used
#   on installs puppeteer itself. I.e the web project's puppeteer dependency will take precedence.
RUN yarn add puppeteer@chrome-$CHROME_VERSION

# It's a good idea to use dumb-init to help prevent zombie chrome processes.
ADD https://github.com/Yelp/dumb-init/releases/download/v1.2.0/dumb-init_1.2.0_amd64 /usr/local/bin/dumb-init
RUN chmod +x /usr/local/bin/dumb-init

ENTRYPOINT ["dumb-init", "--", "yarn"]
