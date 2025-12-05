// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func schemaOptionalClaims() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"name": {
					Description: "The name of the optional claim",
					Type:        pluginsdk.TypeString,
					Required:    true,
				},

				"source": {
					Description: "The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object",
					Type:        pluginsdk.TypeString,
					Optional:    true,
					ValidateFunc: validation.StringInSlice(
						[]string{"user"},
						false,
					),
				},

				"essential": {
					Description: "Whether the claim specified by the client is necessary to ensure a smooth authorization experience",
					Type:        pluginsdk.TypeBool,
					Optional:    true,
					Default:     false,
				},

				"additional_properties": {
					Description: "List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim",
					Type:        pluginsdk.TypeList,
					Optional:    true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
						ValidateFunc: validation.StringInSlice(
							[]string{
								"cloud_displayname",
								"dns_domain_and_sam_account_name",
								"emit_as_roles",
								"include_externally_authenticated_upn_without_hash",
								"include_externally_authenticated_upn",
								"max_size_limit",
								"netbios_domain_and_sam_account_name",
								"on_premise_security_identifier",
								"sam_account_name",
								"use_guid",
							},
							false,
						),
					},
				},
			},
		},
	}
}
