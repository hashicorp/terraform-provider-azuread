package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecurityActivityLog struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	StatusFrom *PartnerSecuritySecurityAlertStatus `json:"statusFrom,omitempty"`
	StatusTo   *PartnerSecuritySecurityAlertStatus `json:"statusTo,omitempty"`

	// The UPN of the partner user who did the status update activity. This attribute is set by the system.
	UpdatedBy *string `json:"updatedBy,omitempty"`

	// The date and time for the status update activity. This attribute is set by the system. The timestamp type represents
	// date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	UpdatedDateTime *string `json:"updatedDateTime,omitempty"`
}
