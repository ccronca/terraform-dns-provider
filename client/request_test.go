package client

import (
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
		uri, err := c.uriForAPI("host", "ip")
		assert.Nil(t, err)
		assert.Equal(t, "https://basedomain.com?host=host&ip_addr=ip", uri)

	})
}
