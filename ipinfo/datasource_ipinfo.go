package ipinfo

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ipinfo/go/v2/ipinfo"
)

var ipInfoSchema = map[string]*schema.Schema{
	"hostname": {
		Type:        schema.TypeString,
		Description: "Hostname",
		Computed:    true,
	},
	"ip": {
		Type:        schema.TypeString,
		Description: "IP address",
		Optional:    true,
	},
}

func datasourceIPInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIPInfoRead,
		Schema:      ipInfoSchema,
	}
}

func datasourceIPInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*ipinfo.Client)

	ipstring := d.Get("ip").(string)

	ip := net.ParseIP(ipstring)
	r, err := client.GetIPInfo(ip)

	if err != nil {
		return diag.FromErr(err)
	}

	if r != nil && r.IP.String() == ipstring {
		d.SetId(r.IP.String())

		if err := d.Set("ip", r.IP.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("hostname", r.Hostname); err != nil {
			return diag.FromErr(err)
		}

		return nil
	}

	return diag.Errorf("IP %s not found", ipstring)
}
