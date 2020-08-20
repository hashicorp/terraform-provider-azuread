package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
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
			appRole["id"] = *v
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

func ApplicationAllOwners(client *graphrbac.ApplicationsClient, ctx context.Context, appId string) ([]string, error) {
	owners, err := client.ListOwnersComplete(ctx, appId)

	if err != nil {
		return nil, fmt.Errorf("listing existing owners for Application with ID %q: %+v", appId, err)
	}

	existingMembers, err := DirectoryObjectListToIDs(owners, ctx)
	if err != nil {
		return nil, fmt.Errorf("getting object IDs of owners for Application with ID %q: %+v", appId, err)
	}

	return existingMembers, nil
}

func ApplicationAddOwner(client *graphrbac.ApplicationsClient, ctx context.Context, appId string, owner string) error {
	ownerGraphURL := fmt.Sprintf("https://graph.windows.net/%s/directoryObjects/%s", client.TenantID, owner)

	properties := graphrbac.AddOwnerParameters{
		URL: &ownerGraphURL,
	}

	if _, err := client.AddOwner(ctx, appId, properties); err != nil {
		return fmt.Errorf("adding owner %q to Application with ID %q: %+v", owner, appId, err)
	}

	return nil
}

func ApplicationAddOwners(client *graphrbac.ApplicationsClient, ctx context.Context, appId string, owner []string) error {
	for _, ownerUuid := range owner {
		err := ApplicationAddOwner(client, ctx, appId, ownerUuid)

		if err != nil {
			return fmt.Errorf("adding owners to Application with ID %q: %+v", appId, err)
		}
	}

	return nil
}

func ApplicationFindByName(ctx context.Context, client *graphrbac.ApplicationsClient, name string) (*graphrbac.Application, error) {
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

func ApplicationCheckNameAvailability(ctx context.Context, client *graphrbac.ApplicationsClient, name string) error {
	existingApp, err := ApplicationFindByName(ctx, client, name)
	if err != nil {
		return err
	}
	if existingApp != nil {
		return fmt.Errorf("existing Application with name %q (AppID: %q) was found and `prevent_duplicate_names` was specified", name, *existingApp.AppID)
	}
	return nil
}

type AppRoleId struct {
	ObjectId string
	RoleId   string
}

func (id AppRoleId) String() string {
	return id.ObjectId + "/" + id.RoleId
}

func AppRoleIdFrom(objectId, roleId string) AppRoleId {
	return AppRoleId{
		ObjectId: objectId,
		RoleId:   roleId,
	}
}

func ParseAppRoleId(id string) (AppRoleId, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		return AppRoleId{}, fmt.Errorf("App Role ID should be in the format {objectId}/{roleId} - but got %q", id)
	}

	if _, err := uuid.ParseUUID(parts[0]); err != nil {
		return AppRoleId{}, fmt.Errorf("Object ID isn't a valid UUID (%q): %+v", parts[0], err)
	}

	if _, err := uuid.ParseUUID(parts[1]); err != nil {
		return AppRoleId{}, fmt.Errorf("Role ID isn't a valid UUID (%q): %+v", parts[2], err)
	}

	return AppRoleId{
		ObjectId: parts[0],
		RoleId:   parts[1],
	}, nil
}

func AppRoleFindByRoleId(app graphrbac.Application, roleId string) (role *graphrbac.AppRole) {
	for _, r := range *app.AppRoles {
		if r.ID == nil {
			continue
		}
		if *r.ID == roleId {
			role = &r
			break
		}
	}
	return role
}

func AppRoleForResource(d *schema.ResourceData) (*graphrbac.AppRole, error) {
	// errors should be handled by the validation
	var roleId string
	if v, ok := d.GetOk("role_id"); ok {
		roleId = v.(string)
	} else {
		rid, err := uuid.GenerateUUID()
		if err != nil {
			return nil, err
		}
		roleId = rid
	}

	allowedMemberTypesRaw := d.Get("allowed_member_types").(*schema.Set).List()
	allowedMemberTypes := make([]string, len(allowedMemberTypesRaw))
	for i, a := range allowedMemberTypesRaw {
		allowedMemberTypes[i] = a.(string)
	}

	appRole := graphrbac.AppRole{
		AllowedMemberTypes: &allowedMemberTypes,
		ID:                 utils.String(roleId),
		Description:        utils.String(d.Get("description").(string)),
		DisplayName:        utils.String(d.Get("display_name").(string)),
		IsEnabled:          utils.Bool(d.Get("is_enabled").(bool)),
	}

	if v, ok := d.GetOk("value"); ok {
		appRole.Value = utils.String(v.(string))
	}

	return &appRole, nil
}

func AppRoleAdd(roles *[]graphrbac.AppRole, role *graphrbac.AppRole) (*[]graphrbac.AppRole, error) {
	newRoles := make([]graphrbac.AppRole, len(*roles)+1)
	newRoles[0] = *role

	for i, v := range *roles {
		if *v.ID == *role.ID {
			return nil, fmt.Errorf("App Role with ID %q already exists", *role.ID)
		}
		newRoles[i+1] = v
	}

	return &newRoles, nil
}

func AppRoleUpdate(roles *[]graphrbac.AppRole, role *graphrbac.AppRole) *[]graphrbac.AppRole {
	newRoles := make([]graphrbac.AppRole, len(*roles))

	for i, v := range *roles {
		if *v.ID == *role.ID {
			newRoles[i] = *role
			continue
		}
		newRoles[i] = v
	}

	return &newRoles
}

func AppRoleResultDisableById(existing *[]graphrbac.AppRole, roleId string) *[]graphrbac.AppRole {
	newRoles := make([]graphrbac.AppRole, len(*existing))

	for i, v := range *existing {
		if *v.ID == roleId {
			v.IsEnabled = utils.Bool(false)
		}
		newRoles[i] = v
	}

	return &newRoles
}

func AppRoleResultRemoveById(existing *[]graphrbac.AppRole, roleId string) *[]graphrbac.AppRole {
	newRoles := make([]graphrbac.AppRole, 0)

	for _, v := range *existing {
		if v.ID == nil {
			continue
		}
		if *v.ID == roleId {
			continue
		}
		newRoles = append(newRoles, v)
	}

	return &newRoles
}
