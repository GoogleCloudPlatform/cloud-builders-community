FROM gcr.io/cloud-builders/go:debian

# Set default GOPATH 
ENV GOPATH /go

# Install go dep
RUN go get -u github.com/golang/dep/cmd/dep
ENV PATH=$GOPATH/bin:$PATH

# Install the dep bash script
COPY dep.bash /builder/bin/
RUN chmod +x /builder/bin/dep.bash
ENV PATH=/builder/bin:$PATH

# Build the go_workspace for detecting the workspace
RUN go build -o /builder/go_workspace /builder/go_workspace.go

ENTRYPOINT ["dep.bash"]