package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionStatusDetails struct {
	// The end time of the connection health check. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EndDateTime *string `json:"endDateTime,omitempty"`

	// All checks that are done on the connection.
	HealthChecks *[]CloudPCOnPremisesConnectionHealthCheck `json:"healthChecks,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The start time of the connection health check. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime *string `json:"startDateTime,omitempty"`
}
