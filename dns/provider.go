package dns

import (
	"log"

	"github.com/camilocot/terraform-dns-provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_USER", nil),
				Description: "API User",
			},
			"api_password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_PASSWORD", nil),
				Description: "API Password",
			},
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_URL", nil),
				Description: "API Url",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"camilocot_dns": resourceDns(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	user := d.Get("api_user").(string)
	password := d.Get("api_password").(string)

	c := client.NewClient(user, password)

	if apiURL := d.Get("api_url").(string); apiURL != "" {
		c.SetBaseUrl(apiURL)
	}
	log.Printf("[INFO] API Client successfully initialized.")

	return c, nil
}
