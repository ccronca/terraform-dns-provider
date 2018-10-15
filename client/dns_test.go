package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {

	dns := DnsRecord{Address: "ip", HostName: "host"}
	assert.Equal(t, "ip_addr=ip&host=host", dns.Query())
}

func TestSucessDnsCreate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))
	defer server.Close()
	api := NewClient("user", "password", http.DefaultClient)
	api.SetBaseUrl(server.URL)

	dns, err := api.DnsCreate("address", "hostname")
	assert.Nil(t, err)
	assert.Equal(t, "address", dns.Address)
	assert.Equal(t, "hostname", dns.HostName)
}

func TestFailedDnsCreate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	api := NewClient("user", "password", http.DefaultClient)
	api.SetBaseUrl(server.URL)

	_, err := api.DnsCreate("address", "hostname")
	assert.Error(t, err)
}

func TestSucessDnsDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))
	defer server.Close()
	api := NewClient("user", "password", http.DefaultClient)
	api.SetBaseUrl(server.URL)

	dns, err := api.DnsDelete("hostname")
	assert.Nil(t, err)
	assert.Equal(t, "delete", dns.Address)
	assert.Equal(t, "hostname", dns.HostName)
}

func TestFailedDnsDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	api := NewClient("user", "password", http.DefaultClient)
	api.SetBaseUrl(server.URL)

	_, err := api.DnsDelete("hostname")
	assert.Error(t, err)
}
