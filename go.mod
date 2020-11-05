module github.com/form3tech-oss/haproxy-data-plane-api-client

go 1.14

replace github.com/haproxytech/models => github.com/haproxytech/models/v2 v2.1.1-0.20201014070602-acf2baeea09a

require (
	github.com/go-openapi/errors v0.19.7
	github.com/go-openapi/runtime v0.19.22
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.11
	github.com/haproxytech/models v2.1.0+incompatible
	github.com/stretchr/testify v1.6.1
)
