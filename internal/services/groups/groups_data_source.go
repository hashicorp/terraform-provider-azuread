// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
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
				ExactlyOneOf: []string{"display_names", "display_name_prefix", "object_ids", "return_all"},
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
				ExactlyOneOf: []string{"display_names", "display_name_prefix", "object_ids", "return_all"},
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"display_name_prefix": {
				Description:      "Common display name prefix of the groups",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_names", "display_name_prefix", "object_ids", "return_all"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"ignore_missing": {
				Description:   "Ignore missing groups and return groups that were found. The data source will still fail if no groups are found",
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"return_all"},
			},

			"return_all": {
				Description:   "Retrieve all groups with no filter",
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"ignore_missing"},
				ExactlyOneOf:  []string{"display_names", "display_name_prefix", "object_ids", "return_all"},
			},

			"mail_enabled": {
				Description:   "Whether the groups are mail-enabled",
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"object_ids"},
			},

			"security_enabled": {
				Description:   "Whether the groups are security-enabled",
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"object_ids"},
			},
		},
	}
}

func groupsDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	var groups []msgraph.Group
	var expectedCount int
	var ignoreMissing = d.Get("ignore_missing").(bool)
	var returnAll = d.Get("return_all").(bool)
	var displayNamePrefix = d.Get("display_name_prefix").(string)

	var displayNames []interface{}
	if v, ok := d.GetOk("display_names"); ok {
		displayNames = v.([]interface{})
	}

	var filter []string

	if v, ok := d.GetOkExists("mail_enabled"); ok { //nolint:staticcheck // needed to detect unset booleans
		filter = append(filter, fmt.Sprintf("mailEnabled eq %t", v.(bool)))
	}
	if v, ok := d.GetOkExists("security_enabled"); ok { //nolint:staticcheck // needed to detect unset booleans
		filter = append(filter, fmt.Sprintf("securityEnabled eq %t", v.(bool)))
	}

	if returnAll {
		result, _, err := client.List(ctx, odata.Query{Filter: strings.Join(filter, " and ")})
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
	} else if displayNamePrefix != "" {
		query := odata.Query{Filter: fmt.Sprintf("startsWith(displayName, '%s')", displayNamePrefix)}
		result, _, err := client.List(ctx, query)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name_prefix", "No groups found with display name prefix: %q", displayNamePrefix)
		}
		if result == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API response")
		}
		if len(*result) == 0 {
			return tf.ErrorDiagPathF(err, "display_name_prefix", "No groups found with display name prefix: %q", displayNamePrefix)
		}

		groups = append(groups, *result...)
	} else if len(displayNames) > 0 {
		expectedCount = len(displayNames)
		for _, v := range displayNames {
			displayName := v.(string)
			query := odata.Query{Filter: strings.Join(append(filter, fmt.Sprintf("displayName eq '%s'", displayName)), " and ")}
			result, _, err := client.List(ctx, query)
			if err != nil {
				return tf.ErrorDiagPathF(err, "display_names", "No group found with display name: %q", displayName)
			}
			if result == nil {
				return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API response")
			}

			count := len(*result)
			if count > 1 {
				return tf.ErrorDiagPathF(err, "display_names", "More than one group found with display name: %q", displayName)
			} else if count == 0 {
				if ignoreMissing {
					continue
				}
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
					if ignoreMissing {
						continue
					}
					return tf.ErrorDiagPathF(err, "object_id", "No group found with object ID: %q", objectId)
				}
				return tf.ErrorDiagPathF(err, "object_id", "Retrieving group with object ID: %q", objectId)
			}
			if group == nil {
				return tf.ErrorDiagF(errors.New("API returned nil group"), "Bad API response")
			}

			groups = append(groups, *group)
		}
	}

	if !returnAll && !ignoreMissing && displayNamePrefix == "" && len(groups) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("Expected: %d, Actual: %d", expectedCount, len(groups)), "Unexpected number of groups returned")
	}

	newDisplayNames := make([]string, 0)
	newObjectIds := make([]string, 0)
	for _, group := range groups {
		if group.ID() == nil {
			return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API response")
		}
		if group.DisplayName == nil {
			return tf.ErrorDiagF(errors.New("API returned group with nil displayName"), "Bad API response")
		}

		newObjectIds = append(newObjectIds, *group.ID())
		newDisplayNames = append(newDisplayNames, *group.DisplayName)
	}

	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(newDisplayNames, "-"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for names")
	}

	d.SetId("groups#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	tf.Set(d, "object_ids", newObjectIds)
	tf.Set(d, "display_names", newDisplayNames)
	tf.Set(d, "display_name_prefix", displayNamePrefix)

	return nil
}
