package main

import (
	"github.com/camilocot/terraform-dns-provider/dns"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return dns.Provider()
		},
	})
}
