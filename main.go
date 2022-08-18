package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/trentmillar/terraform-provider-github-release/github-release"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: github_release.Provider})
}
