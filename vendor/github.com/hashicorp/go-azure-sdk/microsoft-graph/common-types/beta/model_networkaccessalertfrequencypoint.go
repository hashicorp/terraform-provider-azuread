package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessAlertFrequencyPoint struct {
	HighSeverityCount          *int64 `json:"highSeverityCount,omitempty"`
	InformationalSeverityCount *int64 `json:"informationalSeverityCount,omitempty"`
	LowSeverityCount           *int64 `json:"lowSeverityCount,omitempty"`
	MediumSeverityCount        *int64 `json:"mediumSeverityCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	TimeStampDateTime *string `json:"timeStampDateTime,omitempty"`
}
