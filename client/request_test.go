package client

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUriForApi(t *testing.T) {
	c := Client{
		apiUser:     "sample_api_use",
		apiPassword: "sample_api_password",
		baseUrl:     "https://basedomain.com",
		HttpClient:  &http.Client{},
	}

	t.Run("Get Uri for api", func(t *testing.T) {
		uri, err := c.uriForAPI("host=test&ip_addr=ip")
		assert.Nil(t, err)
		assert.Equal(t, "https://basedomain.com?host=test&ip_addr=ip", uri)

	})
	t.Run("Get basic Auth encoded", func(t *testing.T) {
		enc := basicAuth(c.apiUser, c.apiPassword)
		assert.Equal(t, base64.StdEncoding.EncodeToString([]byte(c.apiUser+":"+c.apiPassword)), enc)
	})
	t.Run("Create a request", func(t *testing.T) {
		req, err := c.createRequest("host=test&ip_addr=ip")
		assert.Nil(t, err)
		assert.Regexp(t, "^Basic ", req.Header.Get("Authorization"))
	})

}
