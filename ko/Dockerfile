ARG GO_VERSION=1.26
FROM golang:${GO_VERSION} AS kosource
ARG KO_GIT_TAG

RUN git clone -b "${KO_GIT_TAG}" https://github.com/google/ko
WORKDIR ko
RUN go mod tidy && go mod vendor && GOOS=linux go build -mod=vendor -ldflags "-s -w -X github.com/google/ko/pkg/commands.Version=${KO_GIT_TAG}" -o /bin/ko main.go

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
