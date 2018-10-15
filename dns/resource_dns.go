package dns

import (
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
	return nil
}

func resourceDnsUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceDnsCreate(d, m)
}

func resourceDnsDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)
	hostname := d.Get("hostname").(string)
	_, err := c.DnsDelete(hostname)
	return err
}
