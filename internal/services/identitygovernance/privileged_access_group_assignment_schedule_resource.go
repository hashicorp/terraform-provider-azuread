// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/manicminer/hamilton/msgraph"
)

type PrivilegedAccessGroupAssignmentScheduleResource struct{}

func (r PrivilegedAccessGroupAssignmentScheduleResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidatePrivilegedAccessGroupScheduleID
}

var _ sdk.Resource = PrivilegedAccessGroupAssignmentScheduleResource{}

func (r PrivilegedAccessGroupAssignmentScheduleResource) ResourceType() string {
	return "azuread_privileged_access_group_assignment_schedule"
}

func (r PrivilegedAccessGroupAssignmentScheduleResource) ModelObject() interface{} {
	return &PrivilegedAccessGroupScheduleModel{}
}

func (r PrivilegedAccessGroupAssignmentScheduleResource) Arguments() map[string]*pluginsdk.Schema {
	return privilegedAccessGroupScheduleArguments()
}

func (r PrivilegedAccessGroupAssignmentScheduleResource) Attributes() map[string]*pluginsdk.Schema {
	return privilegedAccessGroupScheduleAttributes()
}

func (r PrivilegedAccessGroupAssignmentScheduleResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleRequestsClient

			var model PrivilegedAccessGroupScheduleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			schedule, err := buildScheduleRequest(&model, &metadata)
			if err != nil {
				return err
			}

			properties := msgraph.PrivilegedAccessGroupAssignmentScheduleRequest{
				AccessId:      model.AssignmentType,
				PrincipalId:   &model.PrincipalId,
				GroupId:       &model.GroupId,
				Action:        msgraph.PrivilegedAccessGroupActionAdminAssign,
				Justification: &model.Justification,
				ScheduleInfo:  schedule,
			}

			if model.TicketNumber != "" || model.TicketSystem != "" {
				properties.TicketInfo = &msgraph.TicketInfo{
					TicketNumber: &model.TicketNumber,
					TicketSystem: &model.TicketSystem,
				}
			}

			req, _, err := client.Create(ctx, properties)
			if err != nil {
				return fmt.Errorf("Could not create assignment schedule request, %+v", err)
			}

			if req.ID == nil || *req.ID == "" {
				return fmt.Errorf("ID returned for assignment schedule request is nil/empty")
			}

			if req.Status == msgraph.PrivilegedAccessGroupAssignmentStatusFailed {
				return fmt.Errorf("Assignment schedule request is in a failed state")
			}

			id, err := parse.ParsePrivilegedAccessGroupScheduleID(*req.TargetScheduleId)
			if err != nil {
				return err
			}
			metadata.SetID(id)

			return nil
		},
	}
}

func (r PrivilegedAccessGroupAssignmentScheduleResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			cSchedule := metadata.Client.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleClient
			cRequests := metadata.Client.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleRequestsClient

			var request *msgraph.PrivilegedAccessGroupAssignmentScheduleRequest

			id, err := parse.ParsePrivilegedAccessGroupScheduleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model PrivilegedAccessGroupScheduleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			schedule, status, err := cSchedule.Get(ctx, id.ID())
			if err != nil && status != http.StatusNotFound {
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			// Some details are only available on the request which is used for the create/update of the schedule.
			// Schedule requests are never deleted. New ones are created when changes are made.
			// Therefore on a read, we need to find the latest version of the request.
			// This is to cater for changes being made outside of Terraform.
			requests, _, err := cRequests.List(ctx, odata.Query{
				Filter: fmt.Sprintf("groupId eq '%s' and targetScheduleId eq '%s'", id.GroupId, id.ID()),
				OrderBy: odata.OrderBy{
					Field:     "createdDateTime",
					Direction: odata.Descending,
				},
			})
			if err != nil {
				return fmt.Errorf("listing requests: %+v", err)
			}
			if len(*requests) == 0 {
				if status == http.StatusNotFound {
					return metadata.MarkAsGone(id)
				}
			} else {
				request = pointer.To((*requests)[0])

				model.Justification = *request.Justification
				if request.TicketInfo.TicketNumber != nil {
					model.TicketNumber = *request.TicketInfo.TicketNumber
				}
				if request.TicketInfo.TicketSystem != nil {
					model.TicketSystem = *request.TicketInfo.TicketSystem
				}
				if request.ScheduleInfo.Expiration.Duration != nil {
					model.Duration = *request.ScheduleInfo.Expiration.Duration
				}
			}

			// Typically this is because the request has expired
			// So we populate the model with the schedule details
			if status == http.StatusNotFound {
				model.AssignmentType = request.AccessId
				model.ExpirationDate = request.ScheduleInfo.Expiration.EndDateTime.Format(time.RFC3339)
				model.GroupId = *request.GroupId
				model.PermanentAssignment = *request.ScheduleInfo.Expiration.Type == msgraph.ExpirationPatternTypeNoExpiration
				model.PrincipalId = *request.PrincipalId
				model.StartDate = request.ScheduleInfo.StartDateTime.Format(time.RFC3339)
				model.Status = request.Status
			} else {
				model.AssignmentType = schedule.AccessId
				model.ExpirationDate = schedule.ScheduleInfo.Expiration.EndDateTime.Format(time.RFC3339)
				model.GroupId = *schedule.GroupId
				model.PermanentAssignment = *schedule.ScheduleInfo.Expiration.Type == msgraph.ExpirationPatternTypeNoExpiration
				model.PrincipalId = *schedule.PrincipalId
				model.StartDate = schedule.ScheduleInfo.StartDateTime.Format(time.RFC3339)
				model.Status = schedule.Status
			}

			return metadata.Encode(&model)
		},
	}
}

func (r PrivilegedAccessGroupAssignmentScheduleResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleRequestsClient

			var model PrivilegedAccessGroupScheduleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			schedule, err := buildScheduleRequest(&model, &metadata)
			if err != nil {
				return err
			}

			properties := msgraph.PrivilegedAccessGroupAssignmentScheduleRequest{
				AccessId:      model.AssignmentType,
				PrincipalId:   &model.PrincipalId,
				GroupId:       &model.GroupId,
				Action:        msgraph.PrivilegedAccessGroupActionAdminAssign,
				Justification: &model.Justification,
				ScheduleInfo:  schedule,
			}

			if model.TicketNumber != "" || model.TicketSystem != "" {
				properties.TicketInfo = &msgraph.TicketInfo{
					TicketNumber: &model.TicketNumber,
					TicketSystem: &model.TicketSystem,
				}
			}

			req, _, err := client.Create(ctx, properties)
			if err != nil {
				return fmt.Errorf("Could not create assignment schedule request, %+v", err)
			}

			if req.ID == nil || *req.ID == "" {
				return fmt.Errorf("ID returned for assignment schedule request is nil/empty")
			}

			if req.Status == msgraph.PrivilegedAccessGroupAssignmentStatusFailed {
				return fmt.Errorf("Assignment schedule request is in a failed state")
			}

			return nil
		},
	}
}

func (r PrivilegedAccessGroupAssignmentScheduleResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleRequestsClient

			id, err := parse.ParsePrivilegedAccessGroupScheduleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model PrivilegedAccessGroupScheduleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			switch model.Status {
			case msgraph.PrivilegedAccessGroupAssignmentStatusDenied,
				msgraph.PrivilegedAccessGroupAssignmentStatusFailed,
				msgraph.PrivilegedAccessGroupAssignmentStatusGranted,
				msgraph.PrivilegedAccessGroupAssignmentStatusPendingAdminDecision,
				msgraph.PrivilegedAccessGroupAssignmentStatusPendingApproval,
				msgraph.PrivilegedAccessGroupAssignmentStatusPendingProvisioning,
				msgraph.PrivilegedAccessGroupAssignmentStatusPendingScheduledCreation:
				return cancelAssignmentRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusProvisioned,
				msgraph.PrivilegedAccessGroupAssignmentStatusScheduleCreated:
				return revokeAssignmentRequest(ctx, metadata, client, id, &model)
			case msgraph.PrivilegedAccessGroupAssignmentStatusCanceled,
				msgraph.PrivilegedAccessGroupAssignmentStatusRevoked:
				return metadata.MarkAsGone(id)
			}

			return fmt.Errorf("unable to destroy due to unknown status: %s", model.Status)
		},
	}
}

func cancelAssignmentRequest(ctx context.Context, metadata sdk.ResourceMetaData, client *msgraph.PrivilegedAccessGroupAssignmentScheduleRequestsClient, id *parse.PrivilegedAccessGroupScheduleId) error {
	status, err := client.Cancel(ctx, id.ID())
	if err != nil {
		if status == http.StatusNotFound {
			return metadata.MarkAsGone(id)
		}
		return fmt.Errorf("cancelling %s: %+v", id, err)
	}
	return nil
}

func revokeAssignmentRequest(ctx context.Context, metadata sdk.ResourceMetaData, client *msgraph.PrivilegedAccessGroupAssignmentScheduleRequestsClient, id *parse.PrivilegedAccessGroupScheduleId, model *PrivilegedAccessGroupScheduleModel) error {
	result, status, err := client.Create(ctx, msgraph.PrivilegedAccessGroupAssignmentScheduleRequest{
		ID:          pointer.To(id.ID()),
		AccessId:    model.AssignmentType,
		PrincipalId: &model.PrincipalId,
		GroupId:     &model.GroupId,
		Action:      msgraph.PrivilegedAccessGroupActionAdminRemove,
	})
	if err != nil {
		if status == http.StatusNotFound {
			return metadata.MarkAsGone(id)
		}
		return fmt.Errorf("retrieving %s: %+v", id, err)
	}
	if result == nil {
		return fmt.Errorf("retrieving %s: API error, result was nil", id)
	}
	return nil
}
