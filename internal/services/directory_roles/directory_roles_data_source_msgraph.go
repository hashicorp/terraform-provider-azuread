package directory_roles

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers `github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph`
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func directoryRolesDataSourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.MsClient

	var dirRoles []msgraph.DirectoryRole
	var expectedCount int

	var displayNames []interface{}
	if v, ok := d.GetOk("display_names"); ok {
		displayNames = v.([]interface{})
	}

	if len(displayNames) > 0 {
		expectedCount = len(displayNames)
		for _, v := range displayNames {
			displayName := v.(string)
			listDirRoles, err := helpers.DirectoryRoleFindByName(ctx, client, displayName)
			if err != nil {
				return tf.ErrorDiagPathF(err, "display_name", "Retrieving directory roles (%s)", displayName)
			}

			count := len(listDirRoles)
			if count > 1 {
				return tf.ErrorDiagPathF(err, "display_name", "More than one directory role found with display name: %q", displayName)
			} else if count == 0 {
				return tf.ErrorDiagPathF(err, "display_name", "No directory role found with display name: %q", displayName)
			}

			dirRoles = append(dirRoles, *listDirRoles[0])
		}
	} else if objectIds, ok := d.Get("object_ids").([]interface{}); ok && len(objectIds) > 0 {
		expectedCount = len(objectIds)
		for _, v := range objectIds {
			objectId := v.(string)
			getDirRole, status, err := client.Get(ctx, objectId)
			if err != nil {
				if status == http.StatusNotFound {
					return tf.ErrorDiagPathF(err, "object_id", "No directory role found with object ID: %q", objectId)
				}
				return tf.ErrorDiagPathF(err, "object_id", "Retrieving directory role with object ID: %q", objectId)
			}

			dirRoles = append(dirRoles, *getDirRole)
		}
	}

	if len(dirRoles) != expectedCount {
		return tf.ErrorDiagF(fmt.Errorf("expected: %d, actual: %d", expectedCount, len(dirRoles)), "Unexpected number of directory roles returned")
	}

	newDisplayNames := make([]string, 0)
	newObjectIds := make([]string, 0)
	for _, dirRole := range dirRoles {
		if dirRole.ID == nil {
			return tf.ErrorDiagF(errors.New("API returned directory role with nil object ID"), "Bad API response")
		}
		if dirRole.DisplayName == nil {
			return tf.ErrorDiagF(errors.New("API returned directory role with nil displayName"), "Bad API response")
		}

		newObjectIds = append(newObjectIds, *dirRole.ID)
		newDisplayNames = append(newDisplayNames, *dirRole.DisplayName)
	}

	h := sha1.New()
	if _, err := h.Write([]byte(strings.Join(newDisplayNames, "-"))); err != nil {
		return tf.ErrorDiagF(err, "Unable to compute hash for names")
	}

	d.SetId("directory_roles#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))

	tf.Set(d, "object_ids", newObjectIds)
	tf.Set(d, "display_names", newDisplayNames)

	return nil
}
