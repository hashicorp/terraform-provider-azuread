// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func ResourceUserInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"user_principal_name": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"display_name": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"account_enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  true,
			},

			"age_group": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"business_phones": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"city": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"company_name": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"consent_provided_for_minor": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"cost_center": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"country": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"department": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"division": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"employee_id": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"employee_type": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"force_password_change": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  false,
			},

			"given_name": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"fax_number": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"job_title": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"mail": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				Computed: true,
			},

			"mail_nickname": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				Computed: true,
			},

			"manager_id": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"mobile_phone": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"office_location": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"onpremises_immutable_id": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				Computed: true,
			},

			"other_mails": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"password": {
				Type:      pluginsdk.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},

			"disable_strong_password": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_password_expiration": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  false,
			},

			"postal_code": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"preferred_language": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"show_in_address_list": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  true,
			},

			"state": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"street_address": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"surname": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"usage_location": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"about_me": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"creation_type": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"external_user_state": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"im_addresses": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"onpremises_distinguished_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_domain_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_sam_account_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_security_identifier": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"onpremises_sync_enabled": {
				Type:     pluginsdk.TypeBool,
				Computed: true,
			},

			"onpremises_user_principal_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"proxy_addresses": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"user_type": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceUserInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId := rawState["id"].(string)
	if _, err := uuid.ParseUUID(oldId); err != nil {
		return rawState, fmt.Errorf("parsing ID for `azuread_user`: %+v", err)
	}

	newId := stable.NewUserID(oldId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
