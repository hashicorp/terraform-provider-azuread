package groups

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func groupDataSourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.MsClient

	var group msgraph.Group
	var displayName string

	if v, ok := d.GetOk("display_name"); ok {
		displayName = v.(string)
	} else if v, ok := d.GetOk("name"); ok {
		displayName = v.(string)
	}

	var mailEnabled, securityEnabled *bool
	if v, exists := d.GetOkExists("mail_enabled"); exists { //nolint:SA1019
		mailEnabled = utils.Bool(v.(bool))
	}
	if v, exists := d.GetOkExists("security_enabled"); exists { //nolint:SA1019
		securityEnabled = utils.Bool(v.(bool))
	}

	if displayName != "" {
		filter := fmt.Sprintf("displayName eq '%s'", displayName)
		if mailEnabled != nil {
			filter = fmt.Sprintf("%s and mailEnabled eq %t", filter, *mailEnabled)
		}
		if securityEnabled != nil {
			filter = fmt.Sprintf("%s and securityEnabled eq %t", filter, *securityEnabled)
		}

		params := []string{fmt.Sprintf("display_name: %q", displayName)}
		if mailEnabled != nil {
			params = append(params, fmt.Sprintf("mail_enabled: %t", *mailEnabled))
		}
		if securityEnabled != nil {
			params = append(params, fmt.Sprintf("security_enabled: %t", *securityEnabled))
		}

		groups, _, err := client.List(ctx, filter)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "No group found matching specified parameters (%s)", strings.Join(params, ", "))
		}

		count := len(*groups)
		if count > 1 {
			return tf.ErrorDiagPathF(err, "name", "More than one group found matching specified parameters (%s)", strings.Join(params, ", "))
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "name", "No group found matching specified parameters (%s)", strings.Join(params, ", "))
		}

		group = (*groups)[0]
	} else if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		g, status, err := client.Get(ctx, objectId)
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "No group found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving group with object ID: %q", objectId)
		}
		if g == nil {
			return tf.ErrorDiagPathF(nil, "object_id", "Group not found with object ID: %q", objectId)
		}

		if mailEnabled != nil && (g.MailEnabled == nil || *g.MailEnabled != *mailEnabled) {
			var actual string
			if g.MailEnabled == nil {
				actual = "nil"
			} else {
				actual = fmt.Sprintf("%t", *g.MailEnabled)
			}
			return tf.ErrorDiagPathF(nil, "mail_enabled", "Group with object ID %q does not have the specified mail_enabled setting (expected: %t, actual: %s)", objectId, *mailEnabled, actual)
		}

		if securityEnabled != nil && (g.SecurityEnabled == nil || *g.SecurityEnabled != *securityEnabled) {
			var actual string
			if g.SecurityEnabled == nil {
				actual = "nil"
			} else {
				actual = fmt.Sprintf("%t", *g.SecurityEnabled)
			}
			return tf.ErrorDiagPathF(nil, "security_enabled", "Group with object ID %q does not have the specified security_enabled setting (expected: %t, actual: %s)", objectId, *securityEnabled, actual)
		}

		group = *g
	}

	if group.ID == nil {
		return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API Response")
	}

	d.SetId(*group.ID)

	tf.Set(d, "description", group.Description)
	tf.Set(d, "display_name", group.DisplayName)
	tf.Set(d, "mail_enabled", group.MailEnabled)
	tf.Set(d, "name", group.DisplayName) // TODO: v2.0 remove this
	tf.Set(d, "object_id", group.ID)
	tf.Set(d, "security_enabled", group.SecurityEnabled)

	members, _, err := client.ListMembers(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve group members for group with object ID: %q", d.Id())
	}
	tf.Set(d, "members", members)

	owners, _, err := client.ListOwners(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve group owners for group with object ID: %q", d.Id())
	}
	tf.Set(d, "owners", owners)

	return nil
}
