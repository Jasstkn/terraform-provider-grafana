package grafana

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceLibraryPanel() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a single library panel by name or uid.",
		ReadContext: dataSourceLibraryPanelRead,
		Schema: cloneResourceSchemaForDatasource(ResourceLibraryPanel(), map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the library panel.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique identifier (UID) of the library panel.",
			},
		}),
	}
}

func dataSourceLibraryPanelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client).gapi
	uid := d.Get("uid").(string)

	// get UID from name if specified
	name := d.Get("name").(string)
	if name != "" {
		panel, err := client.LibraryPanelByName(name)
		if err != nil {
			return diag.FromErr(err)
		}
		uid = panel.UID
	}

	d.SetId(uid)
	ReadLibraryPanel(ctx, d, meta)

	return nil
}
