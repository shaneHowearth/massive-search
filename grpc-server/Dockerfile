############################
ARG GO_VERSION=1.10.8

FROM golang:$GO_VERSION as builder

COPY . massive-search$GOPATH/src/massive-search
WORKDIR $GOPATH/src/massive-search
ADD . $GOPATH/src/massive-search

# Install dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
# Fetch dependencies.
RUN dep ensure --vendor-only
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o /go/bin/server massive-search/grpc-server
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/server /go/bin/server
# Expose port 7777 on the container
EXPOSE 7777
# Run the server binary.
CMD ["/go/bin/server"]
