package namecheap

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Namecheap API Key token/password",
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_TOKEN", nil),
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username of the account associated with the API token",
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_USER", nil),
			},
			"tenant": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username of the account to execute commands against, if different than user",
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_TENANT", nil),
			},
			"client_ip": {
				Type:        schema.TypeString,
				Optional:    true, // This parameter is required by Namecheap, but the value isn't referenced.
				Description: "The IPv4 address sending the API request to the Namecheap API",
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_CLIENT_IP", "127.0.0.1"),
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Namecheap API base API URL",
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_BASE_URL", "https://api.namecheap.com/"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"namecheap_nameservers": resourceNameservers(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"namecheap_domain": dataSourceDomain(),
		},
	}
}
