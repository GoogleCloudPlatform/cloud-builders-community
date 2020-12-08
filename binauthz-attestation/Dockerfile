FROM gcr.io/google.com/cloudsdktool/cloud-sdk:alpine

RUN gcloud components install beta --quiet

RUN apk update && apk upgrade --no-cache

RUN apk add --no-cache docker gnupg bash python3

WORKDIR /work
ADD create_binauthz_attestation.sh /work
ADD parse_arguments.py /work
RUN chmod +x /work/create_binauthz_attestation.sh

ENTRYPOINT [ "/work/create_binauthz_attestation.sh" ]
