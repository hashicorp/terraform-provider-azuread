// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
)

func directoryRolesDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: directoryRolesDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_ids": {
				Description: "The object IDs of the roles",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"template_ids": {
				Description: "The template IDs of the roles",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"roles": {
				Description: "A list of roles",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"display_name": {
							Description: "The display name of the directory role",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"template_id": {
							Description: "The object ID of the template associated with the directory role",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"description": {
							Description: "The description of the directory role",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"object_id": {
							Description: "The object ID of the directory role",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func directoryRolesDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient

	directoryRoles, _, err := client.List(ctx)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve roles")
	}
	if directoryRoles == nil {
		return tf.ErrorDiagF(errors.New("API error: nil directoryRoles were returned"), "Retrieving all directory roles")
	}

	objectIds := make([]string, 0)
	templateIds := make([]string, 0)
	roleList := make([]map[string]interface{}, 0)

	for _, r := range *directoryRoles {
		objectIds = append(objectIds, *r.ID())
		templateIds = append(templateIds, *r.RoleTemplateId)

		role := make(map[string]interface{})
		role["description"] = r.Description
		role["display_name"] = r.DisplayName
		role["object_id"] = r.ID()
		role["template_id"] = r.RoleTemplateId
		roleList = append(roleList, role)
	}

	// Generate a unique ID based on result
	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(objectIds, "/"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for Object IDs")
	}

	d.SetId("roles#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	tf.Set(d, "roles", roleList)
	tf.Set(d, "object_ids", objectIds)
	tf.Set(d, "template_ids", templateIds)

	return nil
}
