FROM gcr.io/cloud-builders/git

RUN apt-get update -y \
    && apt-get install -y --no-install-recommends \
    git-lfs \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
