package main

import (
	"github.com/turbot/steampipe-plugin-office365/microsoft365"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: microsoft365.Plugin})
}
