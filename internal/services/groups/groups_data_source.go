package groups

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func groupsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: groupsDataSourceRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"object_ids": {
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_names", "names", "object_ids"},
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"display_names": {
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_names", "names", "object_ids"},
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"names": {
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_names", "names", "object_ids"},
				Deprecated:   "This property has been renamed to `display_names` and will be removed in v2.0 of this provider.",
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},
		},
	}
}

func groupsDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.AadClient

	var groups []graphrbac.ADGroup
	expectedCount := 0

	var names []interface{}
	if v, ok := d.GetOk("display_names"); ok {
		names = v.([]interface{})
	} else if v, ok := d.GetOk("names"); ok {
		names = v.([]interface{})
	}

	if len(names) > 0 {
		expectedCount = len(names)
		for _, name := range names {
			g, err := aadgraph.GroupGetByDisplayName(ctx, client, name.(string), nil, nil)
			if err != nil {
				return tf.ErrorDiagPathF(err, "name", "No group found with display name: %q", name)
			}
			groups = append(groups, *g)
		}
	} else if objectIds, ok := d.Get("object_ids").([]interface{}); ok && len(objectIds) > 0 {
		expectedCount = len(objectIds)
		for _, v := range objectIds {
			resp, err := client.Get(ctx, v.(string))
			if err != nil {
				if utils.ResponseWasNotFound(resp.Response) {
					return tf.ErrorDiagPathF(nil, "object_id", "No group found with object ID: %q", v)
				}

				return tf.ErrorDiagF(err, "Retrieving group with object ID: %q", v)
			}

			groups = append(groups, resp)
		}
	}

	if len(groups) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("Expected: %d, Actual: %d", expectedCount, len(groups)), "Unexpected number of groups returned")
	}

	newNames := make([]string, 0, len(groups))
	newObjectIds := make([]string, 0, len(groups))
	for _, u := range groups {
		if u.ObjectID == nil || u.DisplayName == nil {
			return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API response")
		}

		newObjectIds = append(newObjectIds, *u.ObjectID)
		newNames = append(newNames, *u.DisplayName)
	}

	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(newNames, "-"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for names")
	}

	d.SetId("groups#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	tf.Set(d, "object_ids", newObjectIds)
	tf.Set(d, "display_names", newNames)
	tf.Set(d, "names", newNames)

	return nil
}
