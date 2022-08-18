package github_release

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

var protoV5ProviderFactories = map[string]func() (*schema.Provider, error){
	"github": func() (*schema.Provider, error) {
		return Provider(), nil
	},
	"github-release": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func init() {
	testAccProvider = Provider() //.(*schema.Provider)
	testAccProviders = map[string]*schema.Provider{
		"github-release": testAccProvider,
	}
	//testAccProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
	//	return map[string]terraform.ResourceProviderFactory{
	//		"github": func() (terraform.ResourceProvider, error) {
	//			p := Provider()
	//			*providers = append(*providers, p.(*schema.Provider))
	//			return p, nil
	//		},
	//	}
	//}
}
