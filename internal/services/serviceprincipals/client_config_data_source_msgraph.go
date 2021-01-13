package serviceprincipals

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func clientConfigDataSourceReadMsGraph(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client)

	objectId := client.Claims.ObjectId
	if client.Claims.ObjectId == "" {
		return tf.ErrorDiagPathF(errors.New("oid claim in access token is empty"), "object_id", "Could not determine object ID of authenticated principal")
	}

	d.SetId(fmt.Sprintf("%s-%s-%s", client.TenantID, client.ClientID, objectId))

	tf.Set(d, "tenant_id", client.TenantID)
	tf.Set(d, "client_id", client.ClientID)
	tf.Set(d, "object_id", objectId)

	return nil
}
