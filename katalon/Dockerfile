FROM katalonstudio/katalon:9.6.0

# Install apt dependencies and the Google Cloud SDK
RUN apt-get update && apt-get install -y apt-transport-https gnupg2 curl && \
    echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - && \
    apt-get update && apt-get install -y google-cloud-sdk kubectl google-cloud-sdk-gke-gcloud-auth-plugin

# Verify that gcloud, kubectl, and gke-gcloud-auth-plugin are installed
RUN gcloud --version && kubectl version --client && gke-gcloud-auth-plugin --version
ENV PATH="$PATH:/usr/local/katalon"

# Set default entrypoint to shell
ENTRYPOINT ["/bin/sh", "-c"]