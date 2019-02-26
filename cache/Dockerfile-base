FROM gcr.io/cloud-builders/gcloud-slim

COPY checksum /bin
COPY save_cache /bin
COPY restore_cache /bin

RUN chmod +x /bin/checksum /bin/save_cache /bin/restore_cache