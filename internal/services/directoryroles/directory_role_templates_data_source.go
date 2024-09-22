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

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroletemplates/stable/directoryroletemplate"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func directoryRoleTemplatesDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: directoryRoleTemplatesDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"object_ids": {
				Description: "The object IDs of the role templates",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"role_templates": {
				Description: "A list of role templates",
				Type:        pluginsdk.TypeList,
				Computed:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"display_name": {
							Description: "The display name of the directory role template",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"description": {
							Description: "The description of the directory role template",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"object_id": {
							Description: "The object ID of the directory role template",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func directoryRoleTemplatesDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleTemplateClient

	resp, err := client.ListDirectoryRoleTemplates(ctx, directoryroletemplate.DefaultListDirectoryRoleTemplatesOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve role templates")
	}

	directoryRoleTemplates := resp.Model
	if directoryRoleTemplates == nil {
		return tf.ErrorDiagF(errors.New("API error: nil directoryRoleTemplates were returned"), "Retrieving all directory role templates")
	}

	objectIds := make([]string, 0)
	templateList := make([]map[string]interface{}, 0)

	for _, r := range *directoryRoleTemplates {
		// Skip the implicit "Users" role as it's non-assignable
		if r.Id == nil || r.DisplayName == nil || r.DisplayName.GetOrZero() == "User" {
			continue
		}

		objectIds = append(objectIds, *r.Id)

		template := make(map[string]interface{})
		template["description"] = r.Description.GetOrZero()
		template["display_name"] = r.DisplayName.GetOrZero()
		template["object_id"] = pointer.From(r.Id)

		templateList = append(templateList, template)
	}

	// Generate a unique ID based on result
	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(objectIds, "/"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for Object IDs")
	}

	d.SetId("templates#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	tf.Set(d, "role_templates", templateList)
	tf.Set(d, "object_ids", objectIds)

	return nil
}
