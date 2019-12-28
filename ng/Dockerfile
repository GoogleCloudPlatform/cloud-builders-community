FROM node:10.16.3

ARG ng_version=latest
RUN npm install -g @angular/cli@$ng_version --unsafe-perms && \
    ng version

ENTRYPOINT ["ng"]
