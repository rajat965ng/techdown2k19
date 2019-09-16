package main

import (
	"github.com/hashicorp/terraform/plugin"
	"./xprovider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: xprovider.Provider,
	})
}