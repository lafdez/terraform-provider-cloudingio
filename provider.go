package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &CloudingIOProvider{}
var _ provider.ProviderWithFunctions = &CloudingIOProvider{}
var _ provider.PorviderWithEphemeralResources = &CloudingIOProvider{}
var _ provider.ProviderWithDataSources = &CloudingIOProvider{}

type CloudingIOProvider struct {
	version string
}

type CloudingIOProviderModel struct {
	APIEndpoint string
	APIKey string
}

func Provider() *schema.Provider {
	return &schema.Provider {
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        types.StringType,
				Required:    true,
			},
			"api_key": {
				Type:        types.StringType,
				Required:    true,
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics {
	config := &CloudingIOProviderModel {
		APIEndpoint: d.Get("endpoint").(string),
		APIKey: d.Get("api_key").(string),
	}
	return config, nil
}

