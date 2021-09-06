package groups

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func groupsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: groupsDataSourceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_ids": {
				Description:  "The object IDs of the groups",
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_names", "object_ids", "return_all"},
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"display_names": {
				Description:  "The display names of the groups",
				Type:         schema.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_names", "object_ids", "return_all"},
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"return_all": {
				Description:  "Retrieve all groups with no filter",
				Type:         schema.TypeBool,
				Optional:     true,
				ExactlyOneOf: []string{"display_names", "object_ids", "return_all"},
			},

			"security_enabled": {
				Description: "Retrieve only groups that are security_enabled groups",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed: true,
			},
		},
	}
}

func groupsDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient
	client.BaseClient.DisableRetries = true

	var groups []msgraph.Group
	var expectedCount int
	var returnAll = d.Get("return_all").(bool)
	var securityEnabled = d.Get("security_enabled").(bool)

	var displayNames []interface{}
	if v, ok := d.GetOk("display_names"); ok {
		displayNames = v.([]interface{})
	}

	if returnAll {
		result, _, err := client.List(ctx, odata.Query{})
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve groups")
		}
		if result == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}
		if len(*result) == 0 {
			return tf.ErrorDiagPathF(err, "return_all", "No groups found")
		}

		groups = append(groups, *result...)
	} else if len(displayNames) > 0 {
		expectedCount = len(displayNames)
		for _, v := range displayNames {
			displayName := v.(string)
			query := odata.Query{
				Filter: fmt.Sprintf("displayName eq '%s'", displayName),
			}
			result, _, err := client.List(ctx, query)
			if err != nil {
				return tf.ErrorDiagPathF(err, "display_names", "No group found with display name: %q", displayName)
			}

			count := len(*result)

			if securityEnabled {
				groups = append(groups, *result...)
			} else if count > 1 {
				return tf.ErrorDiagPathF(err, "display_names", "More than one group found with display name: %q", displayName)
			} else if count == 0 {
				return tf.ErrorDiagPathF(err, "display_names", "No group found with display name: %q", displayName)
			}

			groups = append(groups, (*result)[0])
		}
	} else if objectIds, ok := d.Get("object_ids").([]interface{}); ok && len(objectIds) > 0 {
		expectedCount = len(objectIds)
		for _, v := range objectIds {
			objectId := v.(string)
			group, status, err := client.Get(ctx, objectId, odata.Query{})
			if err != nil {
				if status == http.StatusNotFound {
					return tf.ErrorDiagPathF(err, "object_id", "No group found with object ID: %q", objectId)
				}
				return tf.ErrorDiagPathF(err, "object_id", "Retrieving group with object ID: %q", objectId)
			}

			groups = append(groups, *group)
		}
	}

	if !returnAll && !securityEnabled && len(groups) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("Expected: %d, Actual: %d", expectedCount, len(groups)), "Unexpected number of groups returned")
	}

	newDisplayNames := make([]string, 0)
	newObjectIds := make([]string, 0)
	for _, group := range groups {
		if group.ID == nil {
			return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API response")
		}
		if group.DisplayName == nil {
			return tf.ErrorDiagF(errors.New("API returned group with nil displayName"), "Bad API response")
		}
		if securityEnabled {
			if *group.SecurityEnabled {
				newObjectIds = append(newObjectIds, *group.ID)
				newDisplayNames = append(newDisplayNames, *group.DisplayName)
			}
		} else {
			newObjectIds = append(newObjectIds, *group.ID)
			newDisplayNames = append(newDisplayNames, *group.DisplayName)
		}
	}
	// Check if securityEnabled has caused the number of returned groups to be 0
	if len(newObjectIds) == 0 && securityEnabled {
		return tf.ErrorDiagF(errors.New("No groups found with 'security_enabled = true'"), "Unexpected Number of groups returned")
	}

	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(newDisplayNames, "-"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for names")
	}

	d.SetId("groups#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	tf.Set(d, "object_ids", newObjectIds)
	tf.Set(d, "display_names", newDisplayNames)

	return nil
}
