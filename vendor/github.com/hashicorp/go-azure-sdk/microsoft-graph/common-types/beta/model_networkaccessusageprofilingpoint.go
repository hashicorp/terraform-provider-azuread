package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessUsageProfilingPoint struct {
	InternetAccessTrafficCount     *int64 `json:"internetAccessTrafficCount,omitempty"`
	Microsoft365AccessTrafficCount *int64 `json:"microsoft365AccessTrafficCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PrivateAccessTrafficCount *int64  `json:"privateAccessTrafficCount,omitempty"`
	TimeStampDateTime         *string `json:"timeStampDateTime,omitempty"`
	TotalTrafficCount         *int64  `json:"totalTrafficCount,omitempty"`
}
