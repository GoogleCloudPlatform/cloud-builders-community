FROM gcr.io/cloud-builders/go AS build-env

ADD ./ /go/src/

ENV GOBIN=/go/bin
RUN go get -d -v ./...
RUN go install /go/src/main.go

ENTRYPOINT [ "/go/bin/main" ]
