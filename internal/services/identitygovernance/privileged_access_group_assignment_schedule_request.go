// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type PrivilegedAccessGroupAssignmentScheduleRequestResource struct{}

func (r PrivilegedAccessGroupAssignmentScheduleRequestResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validation.IsUUID
}

var _ sdk.Resource = PrivilegedAccessGroupAssignmentScheduleRequestResource{}

func (r PrivilegedAccessGroupAssignmentScheduleRequestResource) ResourceType() string {
	return "azuread_privileged_access_group_assignment_schedule_request"
}

func (r PrivilegedAccessGroupAssignmentScheduleRequestResource) ModelObject() interface{} {
	return &PrivilegedAccessGroupScheduleRequestModel{}
}

func (r PrivilegedAccessGroupAssignmentScheduleRequestResource) Arguments() map[string]*pluginsdk.Schema {
	return privilegedAccessGroupScheduleRequestArguments()
}

func (r PrivilegedAccessGroupAssignmentScheduleRequestResource) Attributes() map[string]*pluginsdk.Schema {
	return privilegedAccessGroupScheduleRequestAttributes()
}

func (r PrivilegedAccessGroupAssignmentScheduleRequestResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleRequestsClient

			var model PrivilegedAccessGroupScheduleRequestModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			schedule, err := buildRequestSchedule(&model, &metadata)
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

			id := parse.NewPrivilegedAccessGroupAssignmentScheduleRequestID(*req.ID)
			metadata.SetID(id)

			return nil
		},
	}
}

func (r PrivilegedAccessGroupAssignmentScheduleRequestResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleRequestsClient

			id, err := parse.ParsePrivilegedAccessGroupAssignmentScheduleRequestID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model PrivilegedAccessGroupScheduleRequestModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// Schedule requests are never deleted. New ones are created when changes are made.
			// Therefore on a read, we need to find the latest version of the request.
			// This is to cater for changes being made outside of Terraform.
			requests, _, err := client.List(ctx, odata.Query{
				Filter: fmt.Sprintf("groupId eq '%s' and principalId eq '%s'", model.GroupId, model.PrincipalId),
				OrderBy: odata.OrderBy{
					Field:     "createdDateTime",
					Direction: odata.Descending,
				},
			})
			if err != nil {
				return fmt.Errorf("listing requests: %+v", err)
			}
			if len(*requests) == 0 {
				return metadata.MarkAsGone(id)
			}
			request := (*requests)[0]

			if slices.Contains([]string{
				msgraph.PrivilegedAccessGroupAssignmentStatusCanceled,
				msgraph.PrivilegedAccessGroupAssignmentStatusRevoked,
			}, request.Status) {
				metadata.MarkAsGone(id)
			} else {
				model.AssignmentType = request.AccessId
				model.GroupId = *request.GroupId
				model.Justification = *request.Justification
				model.PermanentAssignment = *request.ScheduleInfo.Expiration.Type == msgraph.ExpirationPatternTypeNoExpiration
				model.PrincipalId = *request.PrincipalId
				model.StartDate = request.ScheduleInfo.StartDateTime.Format(time.RFC3339)
				model.Status = request.Status
				model.TargetScheduleId = *request.TargetScheduleId

				if request.ScheduleInfo.Expiration.EndDateTime != nil {
					model.ExpirationDate = request.ScheduleInfo.Expiration.EndDateTime.Format(time.RFC3339)
				}
				if request.ScheduleInfo.Expiration.Duration != nil {
					model.Duration = *request.ScheduleInfo.Expiration.Duration
				}

				if request.TicketInfo.TicketNumber != nil {
					model.TicketNumber = *request.TicketInfo.TicketNumber
				}
				if request.TicketInfo.TicketSystem != nil {
					model.TicketSystem = *request.TicketInfo.TicketSystem
				}

				// Update the ID if it has changed
				if *request.ID != id.ID() {
					id = parse.NewPrivilegedAccessGroupAssignmentScheduleRequestID(*request.ID)
					metadata.SetID(id)
				}
			}

			return metadata.Encode(&model)
		},
	}
}

func (r PrivilegedAccessGroupAssignmentScheduleRequestResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleRequestsClient

			id, err := parse.ParsePrivilegedAccessGroupAssignmentScheduleRequestID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model PrivilegedAccessGroupScheduleRequestModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			switch model.Status {
			case msgraph.PrivilegedAccessGroupAssignmentStatusDenied:
				return cancelAssignmentRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusFailed:
				return cancelAssignmentRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusGranted:
				return cancelAssignmentRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusPendingAdminDecision:
				return cancelAssignmentRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusPendingApproval:
				return cancelAssignmentRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusPendingProvisioning:
				return cancelAssignmentRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusPendingScheduledCreation:
				return cancelAssignmentRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusProvisioned:
				return revokeAssignmentRequest(ctx, metadata, client, id, &model)
			case msgraph.PrivilegedAccessGroupAssignmentStatusScheduleCreated:
				return revokeAssignmentRequest(ctx, metadata, client, id, &model)
			case msgraph.PrivilegedAccessGroupAssignmentStatusCanceled:
				return metadata.MarkAsGone(id)
			case msgraph.PrivilegedAccessGroupAssignmentStatusRevoked:
				return metadata.MarkAsGone(id)
			}

			return fmt.Errorf("unknown status: %s", model.Status)
		},
	}
}

func cancelAssignmentRequest(ctx context.Context, metadata sdk.ResourceMetaData, client *msgraph.PrivilegedAccessGroupAssignmentScheduleRequestsClient, id *parse.PrivilegedAccessGroupAssignmentScheduleRequestId) error {
	status, err := client.Cancel(ctx, id.RequestId)
	if err != nil {
		if status == http.StatusNotFound {
			return metadata.MarkAsGone(id)
		}
		return fmt.Errorf("cancelling %s: %+v", id, err)
	}
	return metadata.MarkAsGone(id)
}

func revokeAssignmentRequest(ctx context.Context, metadata sdk.ResourceMetaData, client *msgraph.PrivilegedAccessGroupAssignmentScheduleRequestsClient, id *parse.PrivilegedAccessGroupAssignmentScheduleRequestId, model *PrivilegedAccessGroupScheduleRequestModel) error {
	result, status, err := client.Create(ctx, msgraph.PrivilegedAccessGroupAssignmentScheduleRequest{
		ID:          &id.RequestId,
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
	return metadata.MarkAsGone(id)
}
