FROM gcr.io/cloud-builders/npm:current

ARG ng_version=latest
RUN npm install -g @angular/cli@$ng_version --unsafe-perms && \
    ng version

ENTRYPOINT ["ng"]
