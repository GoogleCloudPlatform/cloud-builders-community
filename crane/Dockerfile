FROM gcr.io/cloud-builders/go as builder

# Only allow static linking, so that crane can be used in a different
# environment.
ENV CGO_ENABLED=0
RUN go get github.com/google/go-containerregistry/cmd/crane

FROM launcher.gcr.io/google/ubuntu16_04

COPY --from=builder /root/go/bin/crane /builder/bin/crane

ENV PATH=/builder/bin/:$PATH

ENTRYPOINT ["crane"]
