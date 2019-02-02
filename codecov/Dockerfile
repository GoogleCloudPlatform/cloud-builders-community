# By using an existing cloud-builders, this layer will be cached on GCB workers.
FROM gcr.io/cloud-builders/curl

WORKDIR /codecov

# https://github.com/codecov/codecov-bash
RUN curl -s https://codecov.io/bash > codecovbin
RUN chmod +x codecovbin

ENTRYPOINT ["/codecov/codecovbin"]
