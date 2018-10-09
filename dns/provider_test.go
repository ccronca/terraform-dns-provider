package dns

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testDnsProvider *schema.Provider

func init() {
	testDnsProvider = Provider()
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testDnsPreCheck(t *testing.T) {
	if v := os.Getenv("API_USER"); v == "" {
		t.Fatal("API_USER must be set for acceptance tests")
	}
	if v := os.Getenv("API_PASSWORD"); v == "" {
		t.Fatal("API_PASSWORD must be set for acceptance tests")
	}
}
