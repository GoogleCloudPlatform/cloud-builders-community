FROM golang:1.13

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go get github.com/google/ko/cmd/ko

FROM gcr.io/cloud-builders/kubectl

COPY --from=0 /go/bin/ko /
COPY ko.bash /builder/ko.bash

# Install Go
RUN curl https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz -o go1.13.4.linux-amd64.tar.gz \
  && tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz \
  && rm go1.13.4.linux-amd64.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin"
RUN mkdir -p /go/src
ENV GOPATH=/go

RUN gcloud auth configure-docker --quiet

ENTRYPOINT ["/builder/ko.bash"]
