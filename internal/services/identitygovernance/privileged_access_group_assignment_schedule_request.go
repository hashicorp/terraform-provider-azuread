// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
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

			schedule := msgraph.RequestSchedule{}
			schedule.Expiration = &msgraph.ExpirationPattern{}

			if model.ExpirationDate != "" || model.Duration != "" {
				if model.Duration != "" {
					schedule.Expiration.Duration = &model.Duration
					schedule.Expiration.Type = pointer.To(msgraph.ExpirationPatternTypeAfterDuration)
				}

				if model.ExpirationDate != "" {
					endDate, err := time.Parse(time.RFC3339, model.ExpirationDate)
					if err != nil {
						return fmt.Errorf("parsing %s: %+v", model.StartDate, err)
					}
					schedule.Expiration.EndDateTime = &endDate
					schedule.Expiration.Type = pointer.To(msgraph.ExpirationPatternTypeAfterDateTime)
				}
			} else if model.PermanentAssignment {
				schedule.Expiration.Type = pointer.To(msgraph.ExpirationPatternTypeNoExpiration)
			} else {
				return fmt.Errorf("either expiration_date or duration must be set, or permanent_assignment must be true")
			}

			if model.StartDate != "" {
				startDate, err := time.Parse(time.RFC3339, model.StartDate)
				if err != nil {
					return fmt.Errorf("parsing %s: %+v", model.StartDate, err)
				}
				schedule.StartDateTime = &startDate
			} else {
				now := time.Now()
				schedule.StartDateTime = &now
			}

			properties := msgraph.PrivilegedAccessGroupAssignmentScheduleRequest{
				AccessId:      model.AssignmentType,
				PrincipalId:   &model.PrincipalId,
				GroupId:       &model.GroupId,
				Action:        msgraph.PrivilegedAccessGroupActionAdminAssign,
				Justification: &model.Justification,
				ScheduleInfo:  &schedule,
				TicketInfo: &msgraph.TicketInfo{
					TicketNumber: &model.TicketNumber,
					TicketSystem: &model.TicketSystem,
				},
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

			result, status, err := client.Get(ctx, id.ID())
			if err != nil {
				if status == http.StatusNotFound {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			if slices.Contains([]string{
				msgraph.PrivilegedAccessGroupAssignmentStatusCanceled,
				msgraph.PrivilegedAccessGroupAssignmentStatusRevoked,
			}, result.Status) {
				metadata.MarkAsGone(id)
			}

			model.AssignmentType = result.AccessId
			model.GroupId = *result.GroupId
			model.Justification = *result.Justification
			model.PermanentAssignment = *result.ScheduleInfo.Expiration.Type == msgraph.ExpirationPatternTypeNoExpiration
			model.PrincipalId = *result.PrincipalId
			model.StartDate = result.ScheduleInfo.StartDateTime.Format(time.RFC3339)
			model.Status = result.Status
			model.TargetScheduleId = *result.TargetScheduleId

			if result.ScheduleInfo.Expiration.EndDateTime != nil {
				model.ExpirationDate = result.ScheduleInfo.Expiration.EndDateTime.Format(time.RFC3339)
			}
			if result.ScheduleInfo.Expiration.Duration != nil {
				model.Duration = *result.ScheduleInfo.Expiration.Duration
			}

			if result.TicketInfo.TicketNumber != nil {
				model.TicketNumber = *result.TicketInfo.TicketNumber
			}
			if result.TicketInfo.TicketSystem != nil {
				model.TicketSystem = *result.TicketInfo.TicketSystem
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
