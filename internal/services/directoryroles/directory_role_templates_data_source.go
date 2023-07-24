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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func directoryRoleTemplatesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: directoryRoleTemplatesDataSourceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_ids": {
				Description: "The object IDs of the role templates",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"role_templates": {
				Description: "A list of role templates",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_name": {
							Description: "The display name of the directory role template",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": {
							Description: "The description of the directory role template",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"object_id": {
							Description: "The object ID of the directory role template",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func directoryRoleTemplatesDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleTemplatesClient

	directoryRoleTemplates, _, err := client.List(ctx)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve role templates")
	}
	if directoryRoleTemplates == nil {
		return tf.ErrorDiagF(errors.New("API error: nil directoryRoleTemplates were returned"), "Retrieving all directory role templates")
	}

	objectIds := make([]string, 0)
	templateList := make([]map[string]interface{}, 0)

	for _, r := range *directoryRoleTemplates {
		// Skip the implicit "Users" role as it's non-assignable
		if r.DisplayName != "User" {
			continue
		}

		objectIds = append(objectIds, *r.ID())

		template := make(map[string]interface{})
		template["description"] = r.Description
		template["display_name"] = r.DisplayName
		template["object_id"] = r.ID()
		templateList = append(templateList, template)
	}

	// Generate a unique ID based on result
	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(objectIds, "/"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for Object IDs")
	}

	d.SetId("templates#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	tf.Set(d, "templates", templateList)
	tf.Set(d, "object_ids", objectIds)

	return nil
}
