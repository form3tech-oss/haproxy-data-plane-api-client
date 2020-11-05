package client

import (
	"flag"
	"testing"

	"github.com/form3tech-oss/haproxy-data-plane-api-client/client/frontend"
	"github.com/form3tech-oss/haproxy-data-plane-api-client/client/server"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const (
	fooBackendName    = "foo"
	statsFrontendName = "stats"
)

var (
	addr string
	path string
	user string
	pass string
)

func init() {
	flag.StringVar(&addr, "haproxy-data-plane-api-addr", "localhost:5555", "the address at which the haproxy dataplane api can be reached")
	flag.StringVar(&path, "haproxy-data-plane-api-path", "/v2", "the path at which the haproxy dataplane api can be reached")
	flag.StringVar(&user, "haproxy-data-plane-api-user", "dataplane-api", "the username to use when authenticating against the haproxy dataplane api")
	flag.StringVar(&pass, "haproxy-data-plane-api-pass", "dataplane-api", "the password to use when authenticating against the haproxy dataplane api")
}

func TestGetFrontend(t *testing.T) {
	c := client.New(addr, path, []string{})
	c.Debug = true
	h := New(c, strfmt.Default)
	f, err := h.Frontend.GetFrontend(frontend.NewGetFrontendParams().WithName(statsFrontendName), client.BasicAuth(user, pass))
	assert.NoError(t, err)
	assert.Equal(t, statsFrontendName, f.Payload.Data.Name)
	s, err := h.Server.GetServers(server.NewGetServersParams().WithBackend(fooBackendName), client.BasicAuth(user, pass))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(s.Payload.Data))
	assert.Equal(t, "ECDHE-RSA-AES256-GCM-SHA384", s.Payload.Data[0].Ciphers)
}
