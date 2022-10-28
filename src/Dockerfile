FROM golang:1.16 as builder
#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/editory/user_service 
WORKDIR $GOPATH/src/gitlab.udevs.io/editory/user_service
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/user_service /

FROM alpine
COPY --from=builder user_service .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent


CMD ["./user_service"]
