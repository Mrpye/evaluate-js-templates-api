############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

#FROM golang:1.17.2-stretch AS builder
# Install git.
# Git is required for fetching the dependencies.

WORKDIR $GOPATH/src/mypackage/
COPY . ./evaluate-js-templates-api

WORKDIR $GOPATH/src/mypackage/evaluate-js-templates-api/
# Fetch dependencies.
# Using go get.
RUN apk update \                                                                                                                                                                                                                        
  && apk add ca-certificates zip wget tar curl gcc musl-dev linux-headers\                                                                                                                                                                                                      
  && update-ca-certificates

RUN go install
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0  GOARCH=386 GOOS=linux go build -o /go/bin/evaluate-js-templates-api
RUN cd /go/bin/
WORKDIR /go/bin/

############################
# STEP 2 build a small image
############################
#FROM ubuntu 
FROM gcr.io/distroless/static
#FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/evaluate-js-templates-api /go/bin/evaluate-js-templates-api

EXPOSE 8080


LABEL version="1.0.0"
LABEL name="evaluate-js-templates-api"
LABEL maintainer="Andrew Pye"
LABEL description="An API to evaluate javascript code"


ENV WEB_IP=localhost
ENV WEB_PORT=8080

ENTRYPOINT ["/go/bin/evaluate-js-templates-api"]