package http

import (
	"testing"

	"github.com/imroc/req"
	"github.com/stretchr/testify/assert"

	. "github.com/osuthy/micro-test"
)

func TestK8Sからのコネクションを作成(t *testing.T) {
	definition := &HttpServerDefinition{
		config: C{
			"name":     "conName",
			"k8s": C{
				"svc":     "sample-server",
			},
		},
	}
	tc := TC{}
	definition.SetConnectionForK8S(tc, "test-namespace")
	c := tc["conName"].(*Client)
	resp, err := req.Get(c.Url + "/ping")
	p(err)
	assert.Equal(t, resp.String(), "pong")
}
