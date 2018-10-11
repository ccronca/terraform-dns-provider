package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	api := NewClient("user", "password", http.DefaultClient)
	assert.Equal(t, api.apiUser, "user")
	assert.Equal(t, api.apiPassword, "password")
	assert.Empty(t, api.baseUrl)
	assert.NotNil(t, api.HttpClient)
}

func TestBaseUrl(t *testing.T) {
	api := NewClient("user", "password", http.DefaultClient)
	api.SetBaseUrl("baseUrl")
	assert.Equal(t, "baseUrl", api.GetBaseUrl())
}
