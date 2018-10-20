package dns

import (
	"fmt"
	"net"

	"github.com/camilocot/terraform-dns-provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsCreate,
		Read:   resourceDnsRead,
		Update: resourceDnsUpdate,
		Delete: resourceDnsDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (warns []string, errs []error) {
					if ip := net.ParseIP(v.(string)).To4(); ip == nil {
						errs = append(errs, fmt.Errorf("%q must be a valid ip address", k))
					}
					return
				},
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDnsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)
	address := d.Get("address").(string)
	hostname := d.Get("hostname").(string)
	d.SetId(hostname)
	_, err := c.DnsCreate(address, hostname)
	return err
}

func resourceDnsRead(d *schema.ResourceData, m interface{}) error {
	hostname := d.Get("hostname").(string)
	address := d.Get("address").(string)
	ips, err := lookupDns(hostname)
	if err != nil {
		d.SetId("")
	} else if ips[0].String() != address {
		err = d.Set("address", address)
	}

	return err
}

func resourceDnsUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceDnsCreate(d, m)
}

func resourceDnsDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)
	hostname := d.Get("hostname").(string)
	_, err := lookupDns(hostname)
	if err != nil {
		_, err = c.DnsDelete(hostname)
	}
	return err
}

func lookupDns(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}
