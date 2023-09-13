// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryobjects

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func directoryObjectDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: directoryObjectDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_id": {
				Description:      "The object ID of the principal",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},
			"type": {
				Description: "The OData type of the principal",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func directoryObjectDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Users.DirectoryObjectsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	var directoryObject *msgraph.DirectoryObject

	objectId := d.Get("object_id").(string)

	directoryObject, _, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagPathF(nil, "object_id", "Directory Object with ID %q was not found", objectId)
	}
	if directoryObject == nil {
		return tf.ErrorDiagF(fmt.Errorf("nil object returned for directory object with ID: %q", objectId), "Bad API Response")
	}
	if directoryObject.ID() == nil {
		return tf.ErrorDiagF(fmt.Errorf("nil object ID returned for directory object with ID: %q", objectId), "Bad API Response")
	}
	if directoryObject.ODataType == nil {
		return tf.ErrorDiagF(fmt.Errorf("nil OData Type returned for directory object with ID: %q", objectId), "Bad API Response")
	}

	d.SetId(*directoryObject.ID())

	switch *directoryObject.ODataType {
	case odata.TypeUser:
		tf.Set(d, "type", "User")
	case odata.TypeGroup:
		tf.Set(d, "type", "Group")
	case odata.TypeServicePrincipal:
		tf.Set(d, "type", "ServicePrincipal")
	default:
		return diag.Errorf("unknown object type %q returned for directory object with ID: %q", *directoryObject.ODataType, objectId)
	}

	return nil
}
