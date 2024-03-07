// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type PrivilegedAccessGroupScheduleRequestModel struct {
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

func privilegedAccessGroupScheduleRequestArguments() map[string]*pluginsdk.Schema {
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
			ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
				msgraph.PrivilegedAccessGroupRelationshipMember,
				msgraph.PrivilegedAccessGroupRelationshipOwner,
				msgraph.PrivilegedAccessGroupRelationshipUnknown,
			}, false)),
		},

		"start_date": {
			Description:      "The date that this assignment starts, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ForceNew:         true,
			Computed:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.IsRFC3339Time),
		},

		"expiration_date": {
			Description:      "The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ForceNew:         true,
			ConflictsWith:    []string{"duration"},
			ValidateDiagFunc: validation.ValidateDiag(validation.IsRFC3339Time),
		},

		"duration": {
			Description:      "The duration of the assignment, formatted as an ISO8601 duration string (e.g. P3D for 3 days)",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ForceNew:         true,
			ConflictsWith:    []string{"expiration_date"},
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},

		"permanent_assignment": {
			Description: "Is the assignment permanent",
			Type:        pluginsdk.TypeBool,
			Optional:    true,
			ForceNew:    true,
			Computed:    true,
		},

		"justification": {
			Description:      "The justification for the assignment",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ForceNew:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},

		"ticket_number": {
			Description:      "The ticket number authorising the assignment",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ForceNew:         true,
			RequiredWith:     []string{"ticket_system"},
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},

		"ticket_system": {
			Description:      "The ticket system authorising the assignment",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ForceNew:         true,
			RequiredWith:     []string{"ticket_number"},
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},
	}
}

func privilegedAccessGroupScheduleRequestAttributes() map[string]*pluginsdk.Schema {
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
