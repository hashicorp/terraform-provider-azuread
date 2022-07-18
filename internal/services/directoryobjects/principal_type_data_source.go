package directoryobjects

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/odata"
	"time"

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
				Description: "The msgraph type of the principal",
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

	fmt.Println(obj.ODataType)

	return nil
}
