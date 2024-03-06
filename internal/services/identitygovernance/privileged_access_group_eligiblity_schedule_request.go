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

type PrivilegedAccessGroupEligibilityScheduleRequestModel struct {
	AssignmentType      string `tfschema:"assignment_type"`
	Duration            string `tfschema:"duration"`
	ExpirationDate      string `tfschema:"expiration_date"`
	GroupId             string `tfschema:"group_id"`
	Justification       string `tfschema:"justification"`
	PermanentAssignment bool   `tfschema:"permanent_assignment"`
	PrincipalId         string `tfschema:"principal_id"`
	StartDate           string `tfschema:"start_date"`
	Status              string `tfschema:"status"`
	TargetScheduleId    string `tfschema:"target_schedule_id"`
	TicketNumber        string `tfschema:"ticket_number"`
	TicketSystem        string `tfschema:"ticket_system"`
}

type PrivilegedAccessGroupEligibilityScheduleRequestResource struct{}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validation.IsUUID
}

var _ sdk.Resource = PrivilegedAccessGroupEligibilityScheduleRequestResource{}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) ResourceType() string {
	return "azuread_privileged_access_group_eligibility_schedule_request"
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) ModelObject() interface{} {
	return &PrivilegedAccessGroupEligibilityScheduleRequestModel{}
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"group_id": {
			Description:      "The ID of the Group representing the scope of the assignment",
			Type:             pluginsdk.TypeString,
			Required:         true,
			ForceNew:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
		},

		"principal_id": {
			Description:      "The ID of the Principal assigned to the schedule",
			Type:             pluginsdk.TypeString,
			Required:         true,
			ForceNew:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
		},

		"assignment_type": {
			Description: "The ID of the assignment to the group",
			Type:        pluginsdk.TypeString,
			Required:    true,
			ForceNew:    true,
			ValidateFunc: validation.StringInSlice([]string{
				msgraph.PrivilegedAccessGroupRelationshipMember,
				msgraph.PrivilegedAccessGroupRelationshipOwner,
				msgraph.PrivilegedAccessGroupRelationshipUnknown,
			}, false),
		},

		"start_date": {
			Description:  "The date that this assignment starts, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			Computed:     true,
			ValidateFunc: validation.IsRFC3339Time,
		},

		"expiration_date": {
			Description:   "The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)",
			Type:          pluginsdk.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"duration", "permanent_assignment"},
			ValidateFunc:  validation.IsRFC3339Time,
		},

		"duration": {
			Description:   "The duration of the assignment, formatted as an ISO8601 duration string (e.g. P3D for 3 days)",
			Type:          pluginsdk.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"expiration_date", "permanent_assignment"},
			ValidateFunc:  validation.StringIsNotEmpty,
		},

		"permanent_assignment": {
			Description:   "Is the assignment permanent",
			Type:          pluginsdk.TypeBool,
			Optional:      true,
			ForceNew:      true,
			Computed:      true,
			ConflictsWith: []string{"expiration_date", "duration"},
		},

		"justification": {
			Description:  "The justification for the assignment",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"ticket_number": {
			Description:  "The ticket number authorising the assignment",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			RequiredWith: []string{"ticket_system"},
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"ticket_system": {
			Description:  "The ticket system authorising the assignment",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			RequiredWith: []string{"ticket_number"},
			ValidateFunc: validation.StringIsNotEmpty,
		},
	}
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"status": {
			Description: "The status of the Schedule Request",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"target_schedule_id": {
			Description: "The ID of the Schedule targeted by the request",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (r PrivilegedAccessGroupEligibilityScheduleRequestResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupEligibilityScheduleRequestClient

			var model PrivilegedAccessGroupEligibilityScheduleRequestModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			schedule := msgraph.RequestSchedule{}
			schedule.Expiration = &msgraph.ExpirationPattern{}

			if model.ExpirationDate != "" || model.Duration != "" {
				if model.Duration != "" {
					schedule.Expiration.Duration = &model.Duration
					schedule.Expiration.Type = msgraph.ExpirationPatternTypeAfterDuration
				}

				if model.ExpirationDate != "" {
					endDate, err := time.Parse(time.RFC3339, model.ExpirationDate)
					if err != nil {
						return fmt.Errorf("parsing %s: %+v", model.StartDate, err)
					}
					schedule.Expiration.EndDateTime = &endDate
					schedule.Expiration.Type = msgraph.ExpirationPatternTypeAfterDateTime
				}
			} else if model.PermanentAssignment {
				schedule.Expiration.Type = msgraph.ExpirationPatternTypeNoExpiration
			} else {
				schedule.Expiration.Type = msgraph.ExpirationPatternTypeNotSpecified
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

			properties := msgraph.PrivilegedAccessGroupEligibilityScheduleRequest{
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
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupEligibilityScheduleRequestClient

			id, err := parse.ParsePrivilegedAccessGroupEligibilityScheduleRequestID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model PrivilegedAccessGroupEligibilityScheduleRequestModel
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
			model.PermanentAssignment = result.ScheduleInfo.Expiration.Type == msgraph.ExpirationPatternTypeNoExpiration
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
			client := metadata.Client.IdentityGovernance.PrivilegedAccessGroupEligibilityScheduleRequestClient

			id, err := parse.ParsePrivilegedAccessGroupEligibilityScheduleRequestID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model PrivilegedAccessGroupEligibilityScheduleRequestModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			switch model.Status {
			case msgraph.PrivilegedAccessGroupEligibilityStatusDenied:
				return cancelRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusFailed:
				return cancelRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusGranted:
				return cancelRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusPendingAdminDecision:
				return cancelRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusPendingApproval:
				return cancelRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusPendingProvisioning:
				return cancelRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusPendingScheduledCreation:
				return cancelRequest(ctx, metadata, client, id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusProvisioned:
				return revokeRequest(ctx, metadata, client, id, &model)
			case msgraph.PrivilegedAccessGroupEligibilityStatusScheduleCreated:
				return revokeRequest(ctx, metadata, client, id, &model)
			case msgraph.PrivilegedAccessGroupEligibilityStatusCanceled:
				return metadata.MarkAsGone(id)
			case msgraph.PrivilegedAccessGroupEligibilityStatusRevoked:
				return metadata.MarkAsGone(id)
			}

			return fmt.Errorf("unknown status: %s", model.Status)
		},
	}
}

func cancelRequest(ctx context.Context, metadata sdk.ResourceMetaData, client *msgraph.PrivilegedAccessGroupEligibilityScheduleRequestClient, id *parse.PrivilegedAccessGroupEligibilityScheduleRequestId) error {
	status, err := client.Cancel(ctx, id.RequestId)
	if err != nil {
		if status == http.StatusNotFound {
			return metadata.MarkAsGone(id)
		}
		return fmt.Errorf("cancelling %s: %+v", id, err)
	}
	return metadata.MarkAsGone(id)
}

func revokeRequest(ctx context.Context, metadata sdk.ResourceMetaData, client *msgraph.PrivilegedAccessGroupEligibilityScheduleRequestClient, id *parse.PrivilegedAccessGroupEligibilityScheduleRequestId, model *PrivilegedAccessGroupEligibilityScheduleRequestModel) error {
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
