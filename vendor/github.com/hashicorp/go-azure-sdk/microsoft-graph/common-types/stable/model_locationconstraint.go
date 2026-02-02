package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationConstraint struct {
	// The client requests the service to include in the response a meeting location for the meeting. If this is true and
	// all the resources are busy, findMeetingTimes won't return any meeting time suggestions. If this is false and all the
	// resources are busy, findMeetingTimes would still look for meeting times without locations.
	IsRequired nullable.Type[bool] `json:"isRequired,omitempty"`

	// Constraint information for one or more locations that the client requests for the meeting.
	Locations *[]LocationConstraintItem `json:"locations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The client requests the service to suggest one or more meeting locations.
	SuggestLocation nullable.Type[bool] `json:"suggestLocation,omitempty"`
}
