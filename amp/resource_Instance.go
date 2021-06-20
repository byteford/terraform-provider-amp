package amp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceInstance contains the definition of the recourse that terraform creates
func resourceInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInstanceCreate,
		ReadContext:   resourceInstanceRead,
		UpdateContext: resourceInstanceUpdate,
		DeleteContext: resourceInstanceDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"module": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"target_ads_instance": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"new_instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port_number": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"auto_configure": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"provision_setting": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceInstanceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	obj := CreateInstanceObj{
		FriendlyName:      d.Get("friendly_name").(string),
		InstanceName:      d.Get("instance_name").(string),
		Module:            d.Get("module").(string),
		TargetADSInstance: d.Get("target_ads_instance").(string),
		NewInstanceId:     d.Get("new_instance_id").(string),
		PortNumber:        d.Get("port_number").(int32),
		AutoConfigure:     d.Get("auto_configure").(bool),
		ProvisionSettings: d.Get("provision_setting").(map[string]string),
	}
	err := CreateInstance(*c, obj)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(obj.NewInstanceId)
	return resourceInstanceRead(ctx, d, m)
}
func resourceInstanceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	instID := d.Id()

	inst, err := GetInstance(*c, instID)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_id", inst.TargetID); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("friendly_name", inst.FriendlyName); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("instance_name", inst.InstanceName); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("module", inst.Module); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("port", inst.Port); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("running", inst.Running); err != nil {
		diag.FromErr(err)
	}
	var diags diag.Diagnostics
	return diags
}
func resourceInstanceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceInstanceRead(ctx, d, m)
}
func resourceInstanceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
