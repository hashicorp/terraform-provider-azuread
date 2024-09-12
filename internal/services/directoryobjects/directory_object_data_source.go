// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryobjects

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
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
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"type": {
				Description: "The OData type of the principal",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func directoryObjectDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryObjects.DirectoryObjectClient

	id := stable.NewDirectoryObjectID(d.Get("object_id").(string))

	resp, err := client.GetDirectoryObject(ctx, id, directoryobject.DefaultGetDirectoryObjectOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "object_id", "%s was not found", id)
		}
		return tf.ErrorDiagF(nil, "retrieving %s: %v", id, err)
	}

	if resp.Model == nil {
		return tf.ErrorDiagF(fmt.Errorf("nil object returned for %s", id), "Bad API Response")
	}

	directoryObject := resp.Model
	switch {
	case directoryObject.DirectoryObject().Id == nil:
		return tf.ErrorDiagF(fmt.Errorf("nil object ID returned for %s", id), "Bad API Response")
	case directoryObject.DirectoryObject().ODataType == nil:
		return tf.ErrorDiagF(fmt.Errorf("nil OData Type returned for %s", id), "Bad API Response")
	}

	d.SetId(id.ID())
	tf.Set(d, "type", odataType(string(pointer.From(directoryObject.DirectoryObject().ODataType))))

	return nil
}
