package ipinfo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ipinfo/go/ipinfo"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("IPINFO_TOKEN", nil),
				Description: "ipinfo.io Authentication Token",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"ipinfo": datasourceIPInfo(),
		},
	}

	p.ConfigureContextFunc = func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}

		// Get access token by signing up a free account at https://ipinfo.io/signup
		authTransport := ipinfo.AuthTransport{Token: d.Get("api_token").(string)}
		httpClient := authTransport.Client()
		client := ipinfo.NewClient(httpClient)

		return client, nil
	}

	return p
}
