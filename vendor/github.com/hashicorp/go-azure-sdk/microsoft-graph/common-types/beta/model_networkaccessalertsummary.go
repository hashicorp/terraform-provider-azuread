package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessAlertSummary struct {
	AlertType *NetworkaccessAlertType `json:"alertType,omitempty"`
	Count     *int64                  `json:"count,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Severity *NetworkaccessAlertSeverity `json:"severity,omitempty"`
}
