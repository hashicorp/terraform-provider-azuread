package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func SchemaAppRolesComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"allowed_member_types": {
					Type:     schema.TypeSet,
					Computed: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},

				"description": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"display_name": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"is_enabled": {
					Type:     schema.TypeBool,
					Computed: true,
				},

				"value": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func SchemaOauth2PermissionsComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"admin_consent_description": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"admin_consent_display_name": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"is_enabled": {
					Type:     schema.TypeBool,
					Computed: true,
				},

				"type": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"user_consent_description": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"user_consent_display_name": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"value": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func SchemaOptionalClaims() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},

				"source": {
					Type:     schema.TypeString,
					Optional: true,
					ValidateFunc: validation.StringInSlice(
						[]string{"user"},
						false,
					),
				},
				"essential": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"additional_properties": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
						ValidateFunc: validation.StringInSlice(
							[]string{
								"dns_domain_and_sam_account_name",
								"emit_as_roles",
								"netbios_domain_and_sam_account_name",
								"sam_account_name",
							},
							false,
						),
					},
				},
			},
		},
	}
}

func FlattenAppRoles(in *[]graphrbac.AppRole) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	appRoles := make([]map[string]interface{}, len(*in))
	for i, role := range *in {
		appRole := make(map[string]interface{})
		if v := role.ID; v != nil {
			appRole["_id"] = *v
		}
		if v := role.AllowedMemberTypes; v != nil {
			appRole["allowed_member_types"] = *v
		}
		if v := role.Description; v != nil {
			appRole["description"] = *v
		}
		if v := role.DisplayName; v != nil {
			appRole["display_name"] = *v
		}
		if v := role.IsEnabled; v != nil {
			appRole["is_enabled"] = *v
		}
		if v := role.Value; v != nil {
			appRole["value"] = *v
		}
		appRoles[i] = appRole
	}

	return appRoles
}

func FlattenOauth2Permissions(in *[]graphrbac.OAuth2Permission) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0)
	for _, p := range *in {
		permission := make(map[string]interface{})
		if v := p.AdminConsentDescription; v != nil {
			permission["admin_consent_description"] = v
		}
		if v := p.AdminConsentDisplayName; v != nil {
			permission["admin_consent_display_name"] = v
		}
		if v := p.ID; v != nil {
			permission["id"] = v
		}
		if v := p.IsEnabled; v != nil {
			permission["is_enabled"] = *v
		}
		if v := p.Type; v != nil {
			permission["type"] = v
		}
		if v := p.UserConsentDescription; v != nil {
			permission["user_consent_description"] = v
		}
		if v := p.UserConsentDisplayName; v != nil {
			permission["user_consent_display_name"] = v
		}
		if v := p.Value; v != nil {
			permission["value"] = v
		}

		result = append(result, permission)
	}

	return result
}

func ApplicationAllOwners(client *graphrbac.ApplicationsClient, ctx context.Context, groupId string) ([]string, error) {
	owners, err := client.ListOwnersComplete(ctx, groupId)

	if err != nil {
		return nil, fmt.Errorf("Error listing existing applications owners from Azure AD Group with ID %q: %+v", groupId, err)
	}

	existingMembers, err := DirectoryObjectListToIDs(owners, ctx)
	if err != nil {
		return nil, fmt.Errorf("Error getting applications IDs of group owners for Azure AD Group with ID %q: %+v", groupId, err)
	}

	log.Printf("[DEBUG] %d members in Azure AD applications with ID: %q", len(existingMembers), groupId)
	return existingMembers, nil
}

func ApplicationAddOwner(client *graphrbac.ApplicationsClient, ctx context.Context, groupId string, owner string) error {
	ownerGraphURL := fmt.Sprintf("https://graph.windows.net/%s/directoryObjects/%s", client.TenantID, owner)

	properties := graphrbac.AddOwnerParameters{
		URL: &ownerGraphURL,
	}

	log.Printf("[DEBUG] Adding owner with id %q to Azure AD applications with id %q", owner, groupId)
	if _, err := client.AddOwner(ctx, groupId, properties); err != nil {
		return fmt.Errorf("Error adding owner %q to Azure AD applications with ID %q: %+v", owner, groupId, err)
	}

	return nil
}

func ApplicationAddOwners(client *graphrbac.ApplicationsClient, ctx context.Context, groupId string, owner []string) error {
	for _, ownerUuid := range owner {
		err := ApplicationAddOwner(client, ctx, groupId, ownerUuid)

		if err != nil {
			return fmt.Errorf("Error while adding owners to Azure AD applications with ID %q: %+v", groupId, err)
		}
	}

	return nil
}

func ApplicationFindByName(client *graphrbac.ApplicationsClient, ctx context.Context, name string) (*graphrbac.Application, error) {
	nameFilter := fmt.Sprintf("displayName eq '%s'", name)
	resp, err := client.List(ctx, nameFilter)

	if err != nil {
		return nil, fmt.Errorf("unable to list Applications with filter %q: %+v", nameFilter, err)
	}

	for _, app := range resp.Values() {
		if *app.DisplayName == name {
			return &app, nil
		}
	}

	return nil, nil
}

func ApplicationCheckNameAvailability(client *graphrbac.ApplicationsClient, ctx context.Context, name string) error {
	existingApp, err := ApplicationFindByName(client, ctx, name)
	if err != nil {
		return err
	}
	if existingApp != nil {
		return fmt.Errorf("existing Application with name %q (AppID: %q) was found and `prevent_duplicate_names` was specified", name, *existingApp.AppID)
	}
	return nil
}
