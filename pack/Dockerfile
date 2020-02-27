FROM busybox
ENV PACK_VERSION=v0.8.1
RUN wget -O- https://github.com/buildpacks/pack/releases/download/${PACK_VERSION}/pack-${PACK_VERSION}-linux.tgz | tar zx

FROM gcr.io/distroless/base
ENTRYPOINT ["/pack"]
COPY --from=0 /pack /
