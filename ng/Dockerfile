FROM gcr.io/cloud-builders/npm

ARG ng_version=latest
RUN npm install -g @angular/cli@$ng_version --unsafe-perms

ENTRYPOINT ["ng"]
