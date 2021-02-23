FROM golang:1.15

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go get github.com/google/ko/cmd/ko

FROM gcr.io/cloud-builders/kubectl

COPY --from=0 /go/bin/ko /
COPY ko.bash /builder/ko.bash

# Install Go
COPY --from=0 /usr/local/go /usr/local/go
ENV PATH="/go/bin:${PATH}:/usr/local/go/bin"
RUN mkdir -p /go/src
ENV GOPATH=/go

RUN gcloud auth configure-docker --quiet

ENTRYPOINT ["/builder/ko.bash"]
