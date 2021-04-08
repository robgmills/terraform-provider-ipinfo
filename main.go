package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/robgmills/terraform-provider-ipinfo/ipinfo"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ipinfo.Provider})
}
