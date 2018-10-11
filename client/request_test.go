package client

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
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

func TestDoFailedRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something bad happened!"))
	}))

	defer server.Close()

	api := NewClient("user", "password", server.Client())
	api.SetBaseUrl(server.URL)

	err := api.doRequest("query")

	assert.NotNil(t, err)
	assert.Error(t, err, "API error 500 Internal Server Error: Something bad happened!")
}

func TestDoSuccessRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	defer server.Close()

	api := NewClient("user", "password", server.Client())
	api.SetBaseUrl(server.URL)

	err := api.doRequest("query")

	assert.Nil(t, err)
}

func TestCreateRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	defer server.Close()

	api := NewClient("user", "password", server.Client())
	api.SetBaseUrl(server.URL)

	req, err := api.createRequest("query")

	assert.Nil(t, err)
	assert.Equal(t, "GET", req.Method)
	assert.Equal(t, http.Header{"Authorization": []string{"Basic " + basicAuth(api.apiUser, api.apiPassword)}}, req.Header)
}
