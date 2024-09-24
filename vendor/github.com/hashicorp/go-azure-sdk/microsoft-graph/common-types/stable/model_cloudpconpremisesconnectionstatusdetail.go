package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionStatusDetail struct {
	// The end time of the connection health check. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time
	// (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z. Read-Only.
	EndDateTime *string `json:"endDateTime,omitempty"`

	// A list of all checks that have been run on the connection. Read-Only.
	HealthChecks *[]CloudPCOnPremisesConnectionHealthCheck `json:"healthChecks,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The start time of the health check. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC).
	// For example, midnight UTC on Jan 1, 2014 appear as 2014-01-01T00:00:00Z. Read-Only.
	StartDateTime *string `json:"startDateTime,omitempty"`
}
