# haproxy-data-plane-api-client

Golang client for the HAProxy Data Plane API.

## Prerequisites

* https://github.com/go-swagger/go-swagger

## Generating

```shell
# Clone https://github.com/haproxytech/models.git.
$ git clone https://github.com/haproxytech/models.git \
    $GOPATH/src/github.com/haproxytech/models
$ cd $GOPATH/src/github.com/haproxytech/models
$ git checkout 404064d
```

```shell
# Clone https://github.com/haproxytech/dataplaneapi-specification.git.
$ git clone ttps://github.com/haproxytech/dataplaneapi-specification.git \
    $GOPATH/src/github.com/haproxytech/dataplaceapi-specification
$ cd $GOPATH/src/github.com/haproxytech/dataplaceapi-specification
$ git checkout 6f4c2f6
```

```shell
# Generate the 'haproxy_spec.yaml' file.
$ cd $GOPATH/src/github.com/haproxytech/dataplaceapi-specification/build
$ go run build.go -file ../haproxy-spec.yaml > haproxy_spec.yaml
```

```shell
# Generate the client.
$ swagger generate client -f haproxy_spec.yaml \
    -A "Data Plane" \
    -t $GOPATH/src/github.com/form3tech-oss/haproxy-data-plane-api-client \
    --existing-models github.com/haproxytech/models \
    --skip-models \
    --tags=Discovery \
    --tags=Information \
    --tags=Specification \
    --tags=Transactions \
    --tags=Sites \
    --tags=Stats \
    --tags=Global \
    --tags=Frontend \
    --tags=Backend \
    --tags=Bind \
    --tags=Server \
    --tags=Configuration \
    --tags=HTTPRequestRule \
    --tags=HTTPResponseRule \
    --tags=BackendSwitchingRule \
    --tags=ServerSwitchingRule \
    --tags=TCPResponseRule \
    --tags=TCPRequestRule \
    --tags=Filter \
    --tags=StickRule \
    --tags=LogTarget \
    --tags=Reloads \
    --tags=ACL \
    --tags=Defaults \
    --tags=StickTable \
    -r ../copyright.txt
```
