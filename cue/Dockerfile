FROM busybox
ENV CUE_VERSION=0.5.0
RUN wget -O- https://github.com/cue-lang/cue/releases/download/v${CUE_VERSION}/cue_v${CUE_VERSION}_linux_amd64.tar.gz | tar zx

FROM gcr.io/distroless/cc
ENTRYPOINT ["/cue"]
COPY --from=0 /cue /
