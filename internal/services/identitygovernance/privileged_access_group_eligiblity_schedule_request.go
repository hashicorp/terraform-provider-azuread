// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type PrivilegedAccessGroupEligibilityScheduleRequestResource struct{}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validation.IsUUID
}

var _ sdk.Resource = PrivilegedAccessGroupEligibilityScheduleRequestResource{}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) ResourceType() string {
	return "azuread_privileged_access_group_eligibility_schedule_request"
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) ModelObject() interface{} {
	return &PrivilegedAccessGroupScheduleRequestModel{}
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) Arguments() map[string]*pluginsdk.Schema {
	return privilegedAccessGroupScheduleRequestArguments()
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) Attributes() map[string]*pluginsdk.Schema {
	return privilegedAccessGroupScheduleRequestAttributes()
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupEligibilityScheduleRequestsClient

			var model PrivilegedAccessGroupScheduleRequestModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			schedule, err := buildRequestSchedule(&model, &metadata)
			if err != nil {
				return err
			}

			properties := msgraph.PrivilegedAccessGroupEligibilityScheduleRequest{
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

			if req.Status == msgraph.PrivilegedAccessGroupEligibilityStatusFailed {
				return fmt.Errorf("Assignment schedule request is in a failed state")
			}

			id := parse.NewPrivilegedAccessGroupEligibilityScheduleRequestID(*req.ID)
			metadata.SetID(id)

			return nil
		},
	}
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupEligibilityScheduleRequestsClient

			id, err := parse.ParsePrivilegedAccessGroupEligibilityScheduleRequestID(metadata.ResourceData.Id())
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
				msgraph.PrivilegedAccessGroupEligibilityStatusCanceled,
				msgraph.PrivilegedAccessGroupEligibilityStatusRevoked,
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

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupEligibilityScheduleRequestsClient

			id, err := parse.ParsePrivilegedAccessGroupEligibilityScheduleRequestID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model PrivilegedAccessGroupScheduleRequestModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			switch model.Status {
			case msgraph.PrivilegedAccessGroupEligibilityStatusDenied:
				return cancelEligibilityRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusFailed:
				return cancelEligibilityRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusGranted:
				return cancelEligibilityRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusPendingAdminDecision:
				return cancelEligibilityRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusPendingApproval:
				return cancelEligibilityRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusPendingProvisioning:
				return cancelEligibilityRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusPendingScheduledCreation:
				return cancelEligibilityRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusProvisioned:
				return revokeEligibilityRequest(ctx, metadata, client, id, &model)
			case msgraph.PrivilegedAccessGroupEligibilityStatusScheduleCreated:
				return revokeEligibilityRequest(ctx, metadata, client, id, &model)
			case msgraph.PrivilegedAccessGroupEligibilityStatusCanceled:
				return metadata.MarkAsGone(id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusRevoked:
				return metadata.MarkAsGone(id)
			}

			return fmt.Errorf("unknown status: %s", model.Status)
		},
	}
}

func cancelEligibilityRequest(ctx context.Context, metadata sdk.ResourceMetaData, client *msgraph.PrivilegedAccessGroupEligibilityScheduleRequestsClient, id *parse.PrivilegedAccessGroupEligibilityScheduleRequestId) error {
	status, err := client.Cancel(ctx, id.RequestId)
	if err != nil {
		if status == http.StatusNotFound {
			return metadata.MarkAsGone(id)
		}
		return fmt.Errorf("cancelling %s: %+v", id, err)
	}
	return metadata.MarkAsGone(id)
}

func revokeEligibilityRequest(ctx context.Context, metadata sdk.ResourceMetaData, client *msgraph.PrivilegedAccessGroupEligibilityScheduleRequestsClient, id *parse.PrivilegedAccessGroupEligibilityScheduleRequestId, model *PrivilegedAccessGroupScheduleRequestModel) error {
	result, status, err := client.Create(ctx, msgraph.PrivilegedAccessGroupEligibilityScheduleRequest{
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
