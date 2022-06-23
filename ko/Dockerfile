FROM golang:1.17 AS kosource
ARG KO_GIT_TAG

RUN GOBIN="/bin/" go install github.com/google/ko@${KO_GIT_TAG}

FROM gcr.io/cloud-builders/kubectl

COPY --from=kosource /bin/ko /
COPY ko.bash /builder/ko.bash

# Install Go
COPY --from=0 /usr/local/go /usr/local/go
ENV PATH="/go/bin:${PATH}:/usr/local/go/bin"
RUN mkdir -p /go/src
ENV GOPATH=/go

RUN gcloud auth configure-docker --quiet

ENTRYPOINT ["/builder/ko.bash"]
