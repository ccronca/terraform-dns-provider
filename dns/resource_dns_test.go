package dns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceDns(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsResource(rName, rAddress, rHostname),
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr("camilocot_dns.dns_server_1", "address", rAddress),
					resource.TestCheckResourceAttr("camilocot_dns.dns_server_1", "hostname", rHostname),
				),
			},
		},
	})
}

func testAccDnsResource(resource, address, hostname string) string {
	return fmt.Sprintf(`
		resource "camilocot_dns" "%s" {
 			address  = %q
			hostname = %q
		}
	`, resource, address, hostname)
}
