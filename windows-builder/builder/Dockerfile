FROM golang AS build-env

ADD ./ /go/src/builder
WORKDIR /go/src/builder

RUN GO111MODULE=on CGO_ENABLED=0 go build -o /go/bin/main

FROM gcr.io/distroless/base-debian10
COPY --from=build-env /go/bin/main /bin/main
ENTRYPOINT [ "/bin/main" ]
