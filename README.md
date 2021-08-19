# haproxy-data-plane-api-client

Golang client for the HAProxy Data Plane API.

## Prerequisites

* https://github.com/go-swagger/go-swagger

## Generating

```shell
# Generate appropriate client swagger spec

# Download the repository with the spec
$ git clone https://github.com/haproxytech/client-native.git \
    $GOPATH/src/github.com/haproxytech/client-native
$ cd $GOPATH/src/github.com/haproxytech/client-native

# Checkout required version - below you can see the version we used to generate the client
$ git checkout 72c9e92462 
$ make models

$ export TARGET_DIR=$GOPATH/src/github.com/form3tech-oss/haproxy-data-plane-api-client

# Copy models to this repo and use them locally
$ rm -rf ${TARGET_DIR}/models && cp -r models ${TARGET_DIR}

# Generate the client

# First get rid of all old files, as swagger only adds/overrides
$ rm -rf ${TARGET_DIR}/client 

# Command below will modify the generated models - that's normal
$ swagger generate client -f specification/build/haproxy_spec.yaml \
    --name "Data Plane" \
    --target $TARGET_DIR \
    --existing-models github.com/form3tech-oss/haproxy-data-plane-api-client/models \
    --copyright-file specification/copyright.txt
```
