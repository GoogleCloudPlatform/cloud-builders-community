FROM redhat/ubi8-minimal:8.8
USER root
RUN microdnf update --refresh --best --noplugins --setopt=install_weak_deps=0 \
 && rm -rf /var/cache/yum \
 && microdnf clean all
RUN microdnf install jq curl
RUN curl -L https://app.harness.io/public/shared/tools/chaos/hce-cli/latest/hce-cli-latest-linux-amd64 --output /usr/local/bin/hce-cli && chmod +x /usr/local/bin/hce-cli
ENTRYPOINT ["hce-cli"]