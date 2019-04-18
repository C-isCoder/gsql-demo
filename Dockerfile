# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/gsql-demo

# Build the tencent-ocr command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install gsql-demo

# Run the tencent-ocr command by default when the container starts.
ENTRYPOINT /go/bin/gsql-demo

# Document that the service listens on port 6663.
EXPOSE 6666