package main

import (
	"log"

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
	d.SetId(address)
	dns, err := c.Create(address, hostname)
	log.Printf("[INFO] %v", dns)
	return err
}

func resourceDnsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDnsUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDnsDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
