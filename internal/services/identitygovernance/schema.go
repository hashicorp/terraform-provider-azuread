package identitygovernance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/odata"
)

func schemaLocalizedContent() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_text": {
				Description: "The default text of this question",
				Type:        schema.TypeString,
				Required:    true,
			},
			"localized_text": {
				Description: "The localized text of the this question",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"language_code": {
							Description:      "The language code of this question content",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validate.ISO639Language,
						},
						"content": {
							Description: "The localized content of this questions",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func schemaUserSet() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subject_type": {
				Description: "Type of users, valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, `externalSponsors`.",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringInSlice([]string{
					odata.ShortTypeExternalSponsors,
					odata.ShortTypeInternalSponsors,
					odata.ShortTypeRequestorManager,
					odata.ShortTypeConnectedOrganizationMembers,
					odata.ShortTypeGroupMembers,
					odata.ShortTypeSingleUser,
				}, true),
			},
			"is_backup": {
				Description: "For a user in an approval stage, this property indicates whether the user is a backup fallback approver.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"object_id": {
				Description: "The ID of the subject.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
