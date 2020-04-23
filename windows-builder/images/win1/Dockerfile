FROM mcr.microsoft.com/windows/servercore:ltsc2019 
ENV CSI_PROXY_PIPE "\\\\.\\pipe\\csipipe"
LABEL description="PD CSI driver"
COPY gce-pd-csi-driver.exe /gce-pd-csi-driver.exe

RUN mkdir C:\tmp

USER ContainerAdministrator
ENTRYPOINT ["/gce-pd-csi-driver.exe"]
