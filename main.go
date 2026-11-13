package main

import (
	"github.com/turbot/steampipe-plugin-microsoft365/microsoft365"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: microsoft365.Plugin})
}
