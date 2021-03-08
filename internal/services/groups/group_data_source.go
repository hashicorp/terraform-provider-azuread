package groups

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func groupDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: groupDataSourceRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "name", "object_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"display_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"display_name", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Deprecated:       "This property has been renamed to `display_name` and will be removed in v2.0 of this provider.",
				ExactlyOneOf:     []string{"display_name", "name", "object_id"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"mail_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"security_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"owners": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func groupDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.AadClient

	var group graphrbac.ADGroup
	var name string

	if v, ok := d.GetOk("display_name"); ok {
		name = v.(string)
	} else if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	var mailEnabled, securityEnabled *bool
	if v, exists := d.GetOkExists("mail_enabled"); exists { //nolint:SA1019
		mailEnabled = utils.Bool(v.(bool))
	}
	if v, exists := d.GetOkExists("security_enabled"); exists { //nolint:SA1019
		securityEnabled = utils.Bool(v.(bool))
	}

	if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		resp, err := client.Get(ctx, objectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return tf.ErrorDiagPathF(nil, "object_id", "No group found with object ID: %q", objectId)
			}

			return tf.ErrorDiagF(err, "Retrieving group with object ID: %q", objectId)
		}

		if mailEnabled != nil && (resp.MailEnabled == nil || *resp.MailEnabled != *mailEnabled) {
			var actual string
			if resp.MailEnabled == nil {
				actual = "nil"
			} else {
				actual = fmt.Sprintf("%t", *resp.MailEnabled)
			}
			return tf.ErrorDiagPathF(nil, "mail_enabled", "Group with object ID %q does not have the specified mail_enabled setting (expected: %t, actual: %s)", objectId, *mailEnabled, actual)
		}

		if securityEnabled != nil && (resp.SecurityEnabled == nil || *resp.SecurityEnabled != *securityEnabled) {
			var actual string
			if resp.SecurityEnabled == nil {
				actual = "nil"
			} else {
				actual = fmt.Sprintf("%t", *resp.SecurityEnabled)
			}
			return tf.ErrorDiagPathF(nil, "security_enabled", "Group with object ID %q does not have the specified security_enabled setting (expected: %t, actual: %s)", objectId, *securityEnabled, actual)
		}

		group = resp
	} else if name != "" {
		g, err := aadgraph.GroupGetByDisplayName(ctx, client, name, mailEnabled, securityEnabled)
		if err != nil {
			params := []string{fmt.Sprintf("display_name: %q", name)}
			if mailEnabled != nil {
				params = append(params, fmt.Sprintf("mail_enabled: %t", *mailEnabled))
			}
			if securityEnabled != nil {
				params = append(params, fmt.Sprintf("security_enabled: %t", *securityEnabled))
			}
			return tf.ErrorDiagPathF(err, "name", "No group found matching specified parameters (%s)", strings.Join(params, ", "))
		}
		group = *g
	}

	if group.ObjectID == nil {
		return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API Response")
	}

	d.SetId(*group.ObjectID)

	tf.Set(d, "object_id", group.ObjectID)
	tf.Set(d, "display_name", group.DisplayName)
	tf.Set(d, "name", group.DisplayName)
	tf.Set(d, "mail_enabled", group.MailEnabled)
	tf.Set(d, "security_enabled", group.SecurityEnabled)

	description := ""
	if v, ok := group.AdditionalProperties["description"]; ok {
		description = v.(string)
	}
	tf.Set(d, "description", description)

	members, err := aadgraph.GroupAllMembers(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve members for group with object ID %q", d.Id())
	}
	tf.Set(d, "members", members)

	owners, err := aadgraph.GroupAllOwners(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for group with object ID %q", d.Id())
	}
	tf.Set(d, "owners", owners)

	return nil
}
