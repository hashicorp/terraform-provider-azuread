package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedPlan struct {
	// The date and time at which the plan was assigned; for example: 2013-01-02T19:32:30Z. The Timestamp type represents
	// date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014
	// is 2014-01-01T00:00:00Z
	AssignedDateTime nullable.Type[string] `json:"assignedDateTime,omitempty"`

	// Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut.
	CapabilityStatus nullable.Type[string] `json:"capabilityStatus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The name of the service; for example, exchange.
	Service nullable.Type[string] `json:"service,omitempty"`

	// A GUID that identifies the service plan. For a complete list of GUIDs and their equivalent friendly service names,
	// see Product names and service plan identifiers for licensing.
	ServicePlanId nullable.Type[string] `json:"servicePlanId,omitempty"`
}
