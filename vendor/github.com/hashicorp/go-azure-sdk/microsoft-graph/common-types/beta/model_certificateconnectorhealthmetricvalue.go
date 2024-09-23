package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateConnectorHealthMetricValue struct {
	// Timestamp for this metric data-point.
	DateTime *string `json:"dateTime,omitempty"`

	// Count of failed requests/operations.
	FailureCount *int64 `json:"failureCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Count of successful requests/operations.
	SuccessCount *int64 `json:"successCount,omitempty"`
}
