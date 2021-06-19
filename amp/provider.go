package amp

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider defines the data terraform uses to make the provider and resources
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AMP_HOST", nil),
				Description: "The URL of the server: eg. `http://127.0.0.1:8123`",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AMP_USERNAME", nil),
				Description: "The username to log in with",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AMP_PASSWORD", nil),
				Description: "The password to log in with",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			//"homeauto_light": resourceLight(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// providerConfigure is used to set up the Client object which is used when calling the API
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	var host string
	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = tempHost
	}

	c, err := NewClient(host, username, password, &http.Client{})
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return c, diags
}
