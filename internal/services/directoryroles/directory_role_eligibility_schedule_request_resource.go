// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	// "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedule"
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
			Read:   pluginsdk.DefaultTimeout(60 * time.Minute),
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
	id := stable.NewRoleManagementDirectoryRoleEligibilityScheduleRequestID(d.Id())

	roleEligibilityScheduleRequest, err := findRoleEligibilityScheduleRequest(ctx, client, d)
	if err != nil {
		return tf.ErrorDiagF(err, "locating %s", id)
	}

	if roleEligibilityScheduleRequest == nil {
		log.Printf("[DEBUG] %s was not found - removing from state", id)
		d.SetId("")
		return nil
	}

	if foundId := roleEligibilityScheduleRequest.Id; foundId != nil && *foundId != id.UnifiedRoleEligibilityScheduleRequestId {
		log.Printf("[DEBUG] %s changed ID, was %s, now %s", id, roleEligibilityScheduleRequest.Id, *foundId)
		d.SetId(*foundId)
	}

	tf.Set(d, "role_definition_id", roleEligibilityScheduleRequest.RoleDefinitionId.GetOrZero())
	tf.Set(d, "principal_id", roleEligibilityScheduleRequest.PrincipalId.GetOrZero())
	tf.Set(d, "justification", roleEligibilityScheduleRequest.Justification.GetOrZero())
	tf.Set(d, "directory_scope_id", roleEligibilityScheduleRequest.DirectoryScopeId.GetOrZero())

	return nil
}

func directoryRoleEligibilityScheduleRequestResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleEligibilityScheduleRequestClient
	id := stable.NewRoleManagementDirectoryRoleEligibilityScheduleRequestID(d.Id())

	resp, err := client.GetDirectoryRoleEligibilityScheduleRequest(ctx, id, directoryroleeligibilityschedulerequest.DefaultGetDirectoryRoleEligibilityScheduleRequestOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	roleEligibilityScheduleRequest := resp.Model
	if roleEligibilityScheduleRequest == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "API Error")
	}

	properties := stable.UnifiedRoleEligibilityScheduleRequest{
		Action:           pointer.To(stable.UnifiedRoleScheduleRequestActions_AdminRemove),
		RoleDefinitionId: roleEligibilityScheduleRequest.RoleDefinitionId,
		PrincipalId:      roleEligibilityScheduleRequest.PrincipalId,
		Justification:    roleEligibilityScheduleRequest.Justification,
		DirectoryScopeId: roleEligibilityScheduleRequest.DirectoryScopeId,
	}

	if _, err = client.CreateDirectoryRoleEligibilityScheduleRequest(ctx, properties, directoryroleeligibilityschedulerequest.DefaultCreateDirectoryRoleEligibilityScheduleRequestOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing role eligibility schedule request %q: %+v", d.Id(), err)
	}

	return nil
}

// func findRoleAssignment(ctx context.Context, client *directoryroleassignment.DirectoryRoleAssignmentClient, roleDefinitionId string, principalId string) (*stable.UnifiedRoleAssignment, error) {
// 	allAssignments, err := client.ListDirectoryRoleAssignmentsComplete(ctx, directoryroleassignment.DefaultListDirectoryRoleAssignmentsOperationOptions())
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	result := &stable.UnifiedRoleAssignment{}
//
// 	for _, v := range allAssignments.Items {
// 		if strings.EqualFold(v.RoleDefinitionId.GetOrZero(), roleDefinitionId) {
// 			if strings.EqualFold(v.PrincipalId.GetOrZero(), principalId) {
// 				return pointer.To(v), nil
// 			}
// 		}
// 	}
//
// 	return result, nil
// }

// func findRoleEligibilitySchedule(ctx context.Context, client *directoryroleeligibilityschedule.DirectoryRoleEligibilityScheduleClient, id stable.RoleManagementDirectoryRoleAssignmentId) (*stable.UnifiedRoleEligibilitySchedule, error) {
// 	schedules, err  := client.ListDirectoryRoleEligibilitySchedulesComplete(ctx, directoryroleeligibilityschedule.DefaultListDirectoryRoleEligibilitySchedulesOperationOptions())
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	for _, schedule := range schedules.Items {
//
// 	}
//
// 	return nil, nil
// }

func findRoleEligibilityScheduleRequest(ctx context.Context, client *directoryroleeligibilityschedulerequest.DirectoryRoleEligibilityScheduleRequestClient, d *pluginsdk.ResourceData) (*stable.UnifiedRoleEligibilityScheduleRequest, error) {
	scheduleRequests, err := client.ListDirectoryRoleEligibilityScheduleRequestsComplete(ctx, directoryroleeligibilityschedulerequest.DefaultListDirectoryRoleEligibilityScheduleRequestsOperationOptions())
	if err != nil {
		return nil, err
	}

	roleDefinitionId := d.Get("role_definition_id").(string)
	principalId := d.Get("principal_id").(string)
	justification := d.Get("justification").(string)
	directoryScopeId := d.Get("directory_scope_id").(string)

	for _, scheduleRequest := range scheduleRequests.Items {
		if strings.EqualFold(scheduleRequest.RoleDefinitionId.GetOrZero(), roleDefinitionId) &&
			strings.EqualFold(scheduleRequest.PrincipalId.GetOrZero(), principalId) &&
			strings.EqualFold(scheduleRequest.Justification.GetOrZero(), justification) &&
			strings.EqualFold(scheduleRequest.DirectoryScopeId.GetOrZero(), directoryScopeId) {
			return pointer.To(scheduleRequest), nil
		}
	}

	return nil, fmt.Errorf("role eligibility schedule request not found, (Role Definition ID: %s / Principal ID: %s / Directory Scope ID: %s)", roleDefinitionId, principalId, directoryScopeId)
}
