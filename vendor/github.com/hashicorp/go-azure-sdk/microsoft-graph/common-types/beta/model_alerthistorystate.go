package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertHistoryState struct {
	// The Application ID of the calling application that submitted an update (PATCH) to the alert. The appId should be
	// extracted from the auth token and not entered manually by the calling application.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// UPN of user the alert was assigned to (note: alert.assignedTo only stores the last value/UPN).
	AssignedTo nullable.Type[string] `json:"assignedTo,omitempty"`

	// Comment entered by signed-in user.
	Comments *[]string `json:"comments,omitempty"`

	// Analyst feedback on the alert in this update. Possible values are: unknown, truePositive, falsePositive,
	// benignPositive.
	Feedback *AlertFeedback `json:"feedback,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Alert status value (if updated). Possible values are: unknown, newAlert, inProgress, resolved, dismissed.
	Status *AlertStatus `json:"status,omitempty"`

	// Date and time of the alert update. The Timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	UpdatedDateTime nullable.Type[string] `json:"updatedDateTime,omitempty"`

	// UPN of the signed-in user that updated the alert (taken from the bearer token - if in user/delegated auth mode).
	User nullable.Type[string] `json:"user,omitempty"`
}
