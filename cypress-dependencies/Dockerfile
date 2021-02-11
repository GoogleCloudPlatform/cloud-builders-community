ARG NODE_VERSION=latest
FROM node:${NODE_VERSION}

RUN apt update && apt install -y libgtk2.0-0 libgtk-3-0 libgbm-dev libnotify-dev libgconf-2-4 libnss3 libxss1 libasound2 libxtst6 xauth xvfb

ENTRYPOINT ["yarn"]
