package dns

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

var (
	API_USER     = os.Getenv("API_USER")
	API_PASSWORD = os.Getenv("API_PASSWORD")
	API_URL      = os.Getenv("API_URL")
)

const (
	rName     = "dns_server_1"
	rAddress  = "1.2.3.4"
	rHostname = "test.com"
)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"camilocot": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if API_USER == "" {
		t.Fatal("API_USER must be set for acceptance tests")
	}
	if API_PASSWORD == "" {
		t.Fatal("API_PASSWORD must be set for acceptance tests")
	}
	if API_URL == "" {
		t.Fatal("API_URL must be set for acceptance tests")
	}
}
