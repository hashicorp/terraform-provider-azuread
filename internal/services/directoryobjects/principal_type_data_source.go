package directoryobjects

import (
	"context"
	"time"

	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func principalTypeDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: principalTypeDataSourceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Description:      "The object ID of the principal",
				Type:             schema.TypeString,
				Optional:         false,
				Computed:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"type": {
				Description: "The OData type of the principal",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func principalTypeDataSourceRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.DirectoryObjectsClient
	client.BaseClient.DisableRetries = true

	objectId := d.Get("object_id").(string)

	obj, _, err := client.Get(context.Background(), objectId, odata.Query{})
	if err != nil {
		return nil
	}

	switch *obj.ODataType {
	// There are many more types to switch here - These are the ones added per the issue
	case odata.TypeUser:
		d.Set("type", "User")
	case odata.TypeSingleUser:
		d.Set("type", "SingleUser")
	case odata.TypeGroup:
		d.Set("type", "Group")
	case odata.TypeServicePrincipal:
		d.Set("type", "ServicePrincipal")
	default:
		return diag.Errorf("unknown principal type returned: %s", *obj.ODataType)
	}

	tf.Set(d, "type", obj.ODataType)

	return nil
}
