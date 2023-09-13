// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	validation2 "github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

func schemaLocalizedContent() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"default_text": {
				Description: "The default text of this question",
				Type:        pluginsdk.TypeString,
				Required:    true,
			},

			"localized_text": {
				Description: "The localized text of this question",
				Type:        pluginsdk.TypeList,
				Optional:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"language_code": {
							Description:      "The language code of this question content",
							Type:             pluginsdk.TypeString,
							Required:         true,
							ValidateDiagFunc: validation2.ISO639Language,
						},

						"content": {
							Description: "The localized content of this question",
							Type:        pluginsdk.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func schemaUserSet() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"subject_type": {
				Description: "Type of users",
				Type:        pluginsdk.TypeString,
				Required:    true,
				ValidateFunc: validation.StringInSlice([]string{
					odata.ShortTypeConnectedOrganizationMembers,
					odata.ShortTypeExternalSponsors,
					odata.ShortTypeGroupMembers,
					odata.ShortTypeInternalSponsors,
					odata.ShortTypeRequestorManager,
					odata.ShortTypeSingleUser,
				}, true),
			},

			"backup": {
				Description: "For a user in an approval stage, this property indicates whether the user is a backup fallback approver",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
			},

			"object_id": {
				Description: "The object ID of the subject",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},
		},
	}
}
