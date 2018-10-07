package main

import (
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
			"recordtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "A",
			},
		},
	}
}

func resourceDnsCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	recordtype := d.Get("recordtype").(string)
	d.SetId(recordtype + "-" + address)
	return nil
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
