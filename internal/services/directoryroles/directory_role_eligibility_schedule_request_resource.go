// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func directoryRoleEligibilityScheduleRequestResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: directoryRoleEligibilityScheduleRequestResourceCreate,
		ReadContext:   directoryRoleEligibilityScheduleRequestResourceRead,
		DeleteContext: directoryRoleEligibilityScheduleRequestResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(10 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
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
				Description:  "The object ID of the directory role for this role eligibility schedule request",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"principal_id": {
				Description:  "The object ID of the member principal",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"directory_scope_id": {
				Description:  "Identifier of the directory object representing the scope of the role eligibility schedule request",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"justification": {
				Description:  "Justification for why the role is assigned",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func directoryRoleEligibilityScheduleRequestResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleEligibilityScheduleRequestClient

	roleDefinitionId := d.Get("role_definition_id").(string)
	principalId := d.Get("principal_id").(string)
	justification := d.Get("justification").(string)
	directoryScopeId := d.Get("directory_scope_id").(string)

	properties := stable.UnifiedRoleEligibilityScheduleRequest{
		Action:           pointer.To(stable.UnifiedRoleScheduleRequestActions_AdminAssign),
		RoleDefinitionId: nullable.Value(roleDefinitionId),
		PrincipalId:      nullable.Value(principalId),
		Justification:    nullable.Value(justification),
		DirectoryScopeId: nullable.Value(directoryScopeId),
		ScheduleInfo: &stable.RequestSchedule{
			StartDateTime: nullable.Value(time.Now().Format(time.RFC3339)),
			Expiration: &stable.ExpirationPattern{
				Type: pointer.To(stable.ExpirationPatternType_NoExpiration),
			},
		},
	}

	options := directoryroleeligibilityschedulerequest.CreateDirectoryRoleEligibilityScheduleRequestOperationOptions{
		RetryFunc: func(resp *http.Response, o *odata.OData) (bool, error) {
			if response.WasNotFound(resp) && o.Error != nil {
				return o.Error.Match("RoleNotFound") || o.Error.Match("SubjectNotFound"), nil
			}
			return false, nil
		},
	}

	resp, err := client.CreateDirectoryRoleEligibilityScheduleRequest(ctx, properties, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating eligibility schedule request for role %q to principal %q: %+v", roleDefinitionId, principalId, err)
	}

	roleEligibilityScheduleRequest := resp.Model
	if roleEligibilityScheduleRequest == nil || roleEligibilityScheduleRequest.Id == nil {
		return tf.ErrorDiagF(errors.New("returned role roleEligibilityScheduleRequest ID was nil"), "API Error")
	}

	id := stable.NewRoleManagementDirectoryRoleEligibilityScheduleRequestID(*roleEligibilityScheduleRequest.Id)
	d.SetId(id.UnifiedRoleEligibilityScheduleRequestId)

	if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetDirectoryRoleEligibilityScheduleRequest(ctx, id, directoryroleeligibilityschedulerequest.DefaultGetDirectoryRoleEligibilityScheduleRequestOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(resp.Model != nil), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for role eligibility schedule request for %q to be created for directory role %q", principalId, roleDefinitionId)
	}

	return directoryRoleEligibilityScheduleRequestResourceRead(ctx, d, meta)
}

func directoryRoleEligibilityScheduleRequestResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleEligibilityScheduleRequestClient
	instanceClient := meta.(*clients.Client).DirectoryRoles.DirectoryRoleEligibilityScheduleInstanceClient
	id := stable.NewRoleManagementDirectoryRoleEligibilityScheduleRequestID(d.Id())

	resp, err := client.GetDirectoryRoleEligibilityScheduleRequest(ctx, id, directoryroleeligibilityschedulerequest.DefaultGetDirectoryRoleEligibilityScheduleRequestOperationOptions())
	if err != nil {
		if !response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagF(err, "Retrieving %s", id)
		}

		// The service typically purges request resources after 45 days while the eligible assignment remains.
		// Request IDs are not a reliable identity for the resulting schedule, so locate the direct assignment by
		// the principal, role definition and scope stored in Terraform state.
		principalId := d.Get("principal_id").(string)
		roleDefinitionId := d.Get("role_definition_id").(string)
		directoryScopeId := d.Get("directory_scope_id").(string)

		roleEligibilityScheduleInstance, err := findDirectRoleEligibilityScheduleInstance(ctx, instanceClient, d.Id(), principalId, roleDefinitionId, directoryScopeId)
		if err != nil {
			return tf.ErrorDiagF(err, "Retrieving role eligibility schedule instances for principal %q", principalId)
		}
		if roleEligibilityScheduleInstance == nil {
			log.Printf("[DEBUG] No direct role eligibility schedule instance was found for %s - removing from state", id)
			d.SetId("")
			return nil
		}

		tf.Set(d, "role_definition_id", roleEligibilityScheduleInstance.RoleDefinitionId.GetOrZero())
		tf.Set(d, "principal_id", roleEligibilityScheduleInstance.PrincipalId.GetOrZero())
		// Schedule instances do not expose justification, so retain the value from configuration/state.
		tf.Set(d, "justification", d.Get("justification").(string))
		tf.Set(d, "directory_scope_id", roleEligibilityScheduleInstance.DirectoryScopeId.GetOrZero())

		return nil
	}

	roleEligibilityScheduleRequest := resp.Model
	if roleEligibilityScheduleRequest == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "API Error")
	}

	tf.Set(d, "role_definition_id", roleEligibilityScheduleRequest.RoleDefinitionId.GetOrZero())
	tf.Set(d, "principal_id", roleEligibilityScheduleRequest.PrincipalId.GetOrZero())
	tf.Set(d, "justification", roleEligibilityScheduleRequest.Justification.GetOrZero())
	tf.Set(d, "directory_scope_id", roleEligibilityScheduleRequest.DirectoryScopeId.GetOrZero())

	return nil
}

func directoryRoleEligibilityScheduleRequestResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleEligibilityScheduleRequestClient

	properties := stable.UnifiedRoleEligibilityScheduleRequest{
		Action:           pointer.To(stable.UnifiedRoleScheduleRequestActions_AdminRemove),
		RoleDefinitionId: nullable.Value(d.Get("role_definition_id").(string)),
		PrincipalId:      nullable.Value(d.Get("principal_id").(string)),
		Justification:    nullable.Value(d.Get("justification").(string)),
		DirectoryScopeId: nullable.Value(d.Get("directory_scope_id").(string)),
	}

	if _, err := client.CreateDirectoryRoleEligibilityScheduleRequest(ctx, properties, directoryroleeligibilityschedulerequest.DefaultCreateDirectoryRoleEligibilityScheduleRequestOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing role eligibility schedule request %q: %+v", d.Id(), err)
	}

	return nil
}

func findDirectRoleEligibilityScheduleInstance(ctx context.Context, client *directoryroleeligibilityscheduleinstance.DirectoryRoleEligibilityScheduleInstanceClient, resourceId, principalId, roleDefinitionId, directoryScopeId string) (*stable.UnifiedRoleEligibilityScheduleInstance, error) {
	options := directoryroleeligibilityscheduleinstance.DefaultListDirectoryRoleEligibilityScheduleInstancesOperationOptions()
	options.Filter = pointer.To(roleEligibilityScheduleInstanceFilter(resourceId, principalId))

	result, err := client.ListDirectoryRoleEligibilityScheduleInstancesComplete(ctx, options)
	if err != nil {
		return nil, err
	}

	return matchDirectRoleEligibilityScheduleInstance(result.Items, resourceId, principalId, roleDefinitionId, directoryScopeId), nil
}

func roleEligibilityScheduleInstanceFilter(resourceId, principalId string) string {
	if principalId == "" {
		return fmt.Sprintf("roleEligibilityScheduleId eq '%s'", odata.EscapeSingleQuote(resourceId))
	}

	return fmt.Sprintf("principalId eq '%s'", odata.EscapeSingleQuote(principalId))
}

func matchDirectRoleEligibilityScheduleInstance(instances []stable.UnifiedRoleEligibilityScheduleInstance, resourceId, principalId, roleDefinitionId, directoryScopeId string) *stable.UnifiedRoleEligibilityScheduleInstance {
	for i := range instances {
		instance := &instances[i]
		if instance.MemberType.GetOrZero() != "Direct" {
			continue
		}

		// During import only the resource ID is available in state.
		if principalId == "" && roleDefinitionId == "" && directoryScopeId == "" {
			if instance.RoleEligibilityScheduleId.GetOrZero() == resourceId {
				return instance
			}
			continue
		}

		if instance.PrincipalId.GetOrZero() == principalId &&
			instance.RoleDefinitionId.GetOrZero() == roleDefinitionId &&
			instance.DirectoryScopeId.GetOrZero() == directoryScopeId {
			return instance
		}
	}

	return nil
}
