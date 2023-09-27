package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func directoryRoleEligibilityScheduleRequestResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: directoryRoleEligibilityScheduleRequestResourceCreate,
		ReadContext:   directoryRoleEligibilityScheduleRequestResourceRead,
		DeleteContext: directoryRoleEligibilityScheduleRequestResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"role_definition_id": {
				Description:      "The object ID of the directory role for this role eligibility schedule request",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"principal_id": {
				Description:      "The object ID of the member principal",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"directory_scope_id": {
				Description:      "Identifier of the directory object representing the scope of the role eligibility schedule request",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"justification": {
				Description:      "Justification for why the role is assigned",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},
		},
	}
}

func directoryRoleEligibilityScheduleRequestResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleEligibilityScheduleRequestClient

	roleDefinitionId := d.Get("role_definition_id").(string)
	principalId := d.Get("principal_id").(string)
	justification := d.Get("justification").(string)
	directoryScopeId := d.Get("directory_scope_id").(string)

	now := time.Now()
	properties := msgraph.UnifiedRoleEligibilityScheduleRequest{
		Action:           utils.String(msgraph.UnifiedRoleScheduleRequestActionAdminAssign),
		RoleDefinitionId: &roleDefinitionId,
		PrincipalId:      &principalId,
		Justification:    &justification,
		DirectoryScopeId: &directoryScopeId,
		ScheduleInfo: &msgraph.RequestSchedule{
			StartDateTime: &now,
			Expiration: &msgraph.ExpirationPattern{
				Type: utils.String(msgraph.ExpirationPatternTypeNoExpiration),
			},
		},
	}

	roleEligibilityScheduleRequest, status, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Eligibility schedule request for role %q to principal %q, received %d with error: %+v", roleDefinitionId, principalId, status, err)
	}
	if roleEligibilityScheduleRequest == nil || roleEligibilityScheduleRequest.ID == nil {
		return tf.ErrorDiagF(errors.New("returned role roleEligibilityScheduleRequest ID was nil"), "API Error")
	}

	d.SetId(*roleEligibilityScheduleRequest.ID)

	if err := helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true

		resr, status, err := client.Get(ctx, *roleEligibilityScheduleRequest.ID, odata.Query{})
		if err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(resr != nil), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for role eligibility schedule request for %q to be created for directory role %q", principalId, roleDefinitionId)
	}

	return directoryRoleEligibilityScheduleRequestResourceRead(ctx, d, meta)
}

func directoryRoleEligibilityScheduleRequestResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleEligibilityScheduleRequestClient

	id := d.Id()
	roleEligibilityScheduleRequest, status, err := client.Get(ctx, id, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] roleEligibilityScheduleRequest with ID %q was not found - removing from state", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving roleEligibilityScheduleRequest %q", id)
	}

	tf.Set(d, "role_definition_id", roleEligibilityScheduleRequest.RoleDefinitionId)
	tf.Set(d, "principal_id", roleEligibilityScheduleRequest.PrincipalId)
	tf.Set(d, "justification", roleEligibilityScheduleRequest.Justification)
	tf.Set(d, "directory_scope_id", roleEligibilityScheduleRequest.DirectoryScopeId)

	return nil
}

func directoryRoleEligibilityScheduleRequestResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleEligibilityScheduleRequestClient

	id := d.Id()
	roleEligibilityScheduleRequest, _, err := client.Get(ctx, id, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving roleEligibilityScheduleRequest %q", id)
	}

	roleEligibilityScheduleRequest.Action = utils.String(msgraph.UnifiedRoleScheduleRequestActionAdminRemove)

	if _, _, err := client.Create(ctx, *roleEligibilityScheduleRequest); err != nil {
		return tf.ErrorDiagF(err, "Deleting role eligibility schedule request %q: %+v", d.Id(), err)
	}
	return nil
}
