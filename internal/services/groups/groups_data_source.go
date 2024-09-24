// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	groupBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func groupsDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: groupsDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_ids": {
				Description:  "The object IDs of the groups",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_names", "display_name_prefix", "object_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},

			"display_names": {
				Description:  "The display names of the groups",
				Type:         pluginsdk.TypeList,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_names", "display_name_prefix", "object_ids", "return_all"},
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"display_name_prefix": {
				Description:  "Common display name prefix of the groups",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"display_names", "display_name_prefix", "object_ids", "return_all"},
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"ignore_missing": {
				Description:   "Ignore missing groups and return groups that were found. The data source will still fail if no groups are found",
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"return_all"},
			},

			"return_all": {
				Description:   "Retrieve all groups with no filter",
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"ignore_missing"},
				ExactlyOneOf:  []string{"display_names", "display_name_prefix", "object_ids", "return_all"},
			},

			"mail_enabled": {
				Description:   "Whether the groups are mail-enabled",
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"object_ids"},
			},

			"security_enabled": {
				Description:   "Whether the groups are security-enabled",
				Type:          pluginsdk.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"object_ids"},
			},
		},
	}
}

func groupsDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta

	var groups []beta.Group
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
		options := groupBeta.ListGroupsOperationOptions{
			Filter: pointer.To(strings.Join(filter, " and ")),
		}
		resp, err := client.ListGroups(ctx, options)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve groups")
		}
		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("model was nil"), "Bad API Response")
		}
		if len(*resp.Model) == 0 {
			return tf.ErrorDiagPathF(err, "return_all", "No groups found")
		}

		groups = append(groups, *resp.Model...)
	} else if displayNamePrefix != "" {
		options := groupBeta.ListGroupsOperationOptions{
			Filter: pointer.To(strings.Join(append(filter, fmt.Sprintf("startsWith(displayName, '%s')", odata.EscapeSingleQuote(displayNamePrefix))), " and ")),
		}
		resp, err := client.ListGroups(ctx, options)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name_prefix", "No groups found with display name prefix: %q", displayNamePrefix)
		}
		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API response")
		}
		if len(*resp.Model) == 0 {
			return tf.ErrorDiagPathF(err, "display_name_prefix", "No groups found with display name prefix: %q", displayNamePrefix)
		}

		groups = append(groups, *resp.Model...)
	} else if len(displayNames) > 0 {
		expectedCount = len(displayNames)
		for _, v := range displayNames {
			displayName := v.(string)
			options := groupBeta.ListGroupsOperationOptions{
				Filter: pointer.To(strings.Join(append(filter, fmt.Sprintf("displayName eq '%s'", odata.EscapeSingleQuote(displayName))), " and ")),
			}
			resp, err := client.ListGroups(ctx, options)
			if err != nil {
				return tf.ErrorDiagPathF(err, "display_names", "No group found with display name: %q", displayName)
			}
			if resp.Model == nil {
				return tf.ErrorDiagF(errors.New("model was nil"), "Bad API response")
			}

			count := len(*resp.Model)
			if count > 1 {
				return tf.ErrorDiagPathF(err, "display_names", "More than one group found with display name: %q", displayName)
			} else if count == 0 {
				if ignoreMissing {
					continue
				}
				return tf.ErrorDiagPathF(err, "display_names", "No group found with display name: %q", displayName)
			}

			groups = append(groups, (*resp.Model)[0])
		}
	} else if objectIds, ok := d.Get("object_ids").([]interface{}); ok && len(objectIds) > 0 {
		expectedCount = len(objectIds)
		for _, v := range objectIds {
			id := beta.NewGroupID(v.(string))
			resp, err := client.GetGroup(ctx, id, groupBeta.DefaultGetGroupOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					if ignoreMissing {
						continue
					}
					return tf.ErrorDiagPathF(err, "object_id", "No group found with object ID: %q", id.GroupId)
				}
				return tf.ErrorDiagPathF(err, "object_id", "Retrieving group with object ID: %q", id.GroupId)
			}
			if resp.Model == nil {
				return tf.ErrorDiagF(errors.New("model was nil"), "Bad API response")
			}

			groups = append(groups, *resp.Model)
		}
	}

	if !returnAll && !ignoreMissing && displayNamePrefix == "" && len(groups) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("expected: %d, actual: %d", expectedCount, len(groups)), "Unexpected number of groups returned")
	}

	newDisplayNames := make([]string, 0)
	newObjectIds := make([]string, 0)
	for _, group := range groups {
		if group.Id == nil {
			return tf.ErrorDiagF(errors.New("group ID was nil"), "Bad API response")
		}
		if group.DisplayName.IsNull() {
			return tf.ErrorDiagF(errors.New("displayName was nil"), "Bad API response")
		}

		newObjectIds = append(newObjectIds, *group.Id)
		newDisplayNames = append(newDisplayNames, group.DisplayName.GetOrZero())
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
