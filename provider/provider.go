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

type CloudingIOProvider struct {
	version string
}

type CloudingIOProviderModel struct {
	APIEndpoint types.String `tfsdk:"endpoint"`
	APIKey      types.String `tfsdk:"api_key"`
}

func (p *CloudingIOProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudingio"
	resp.Version = p.version
}

func (p *CloudingIOProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema {
		Attributes: map[string]schema.Attribute {
			"endpoint": schema.StringAttribute {
				MarkdownDescription: "The API endpoint for Clouding.io",
				Required:            true,
			},
			"api_key": schema.StringAttribute {
				MarkdownDescription: "The API key for Clouding.io",
				Required:    true,
			},
		},
	}
}

func (p *CloudingIOProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CloudingIOProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}



func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudingIOProvider{
			version: version,
		}
	}
}
